/*
   Copyright 2015 Shlomi Noach, courtesy Booking.com

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package process

import (
	"github.com/openark/golib/log"
	"github.com/openark/golib/sqlutils"
	"github.com/openark/orchestrator/go/config"
	"github.com/openark/orchestrator/go/db"
	orcraft "github.com/openark/orchestrator/go/raft"
	"github.com/openark/orchestrator/go/util"
)

/**
尝试把当前的 orchestrator 进程设置成 active_node,
	1、如果设置成功返回 true, nil
	2、如果进程本身就已经是 active_node 也返回 true, nil
	3、如果没能把自己设置成 active_node 就返回 false, nil
*/
// AttemptElection tries to grab leadership (become active node)
func AttemptElection() (bool, error) {
	{
		// 对于 insert ignore 来说在表中没有 ancher = 1 的数据时能插入成功；
		// 对应到业务逻辑上就是说当前 orchestartor 进得是第一个启动的 orchestrator， 这个时候第一个启动的进程会成功 active_node
		sqlResult, err := db.ExecOrchestrator(`
		insert ignore into active_node (
				anchor, hostname, token, first_seen_active, last_seen_active
			) values (
				1, ?, ?, now(), now()
			)
		`,
			ThisHostname, util.ProcessToken.Hash,
		)
		if err != nil {
			return false, log.Errore(err)
		}
		rows, err := sqlResult.RowsAffected()
		if err != nil {
			return false, log.Errore(err)
		}
		if rows > 0 {
			// We managed to insert a row
			return true, nil
		}
	}
	{
		// takeover from a node that has been inactive
		sqlResult, err := db.ExecOrchestrator(`
			update active_node set
				hostname = ?,
				token = ?,
				first_seen_active=now(),
				last_seen_active=now()
			where
				anchor = 1
			  and last_seen_active < (now() - interval ? second)
		`,
			ThisHostname, util.ProcessToken.Hash, config.ActiveNodeExpireSeconds,
		)
		if err != nil {
			return false, log.Errore(err)
		}
		rows, err := sqlResult.RowsAffected()
		if err != nil {
			return false, log.Errore(err)
		}
		if rows > 0 {
			// We managed to update a row: overtaking a previous leader
			return true, nil
		}
	}
	{
		// Update last_seen_active is this very node is already the active node
		sqlResult, err := db.ExecOrchestrator(`
			update active_node set
				last_seen_active=now()
			where
				anchor = 1
				and hostname = ?
				and token = ?
		`,
			ThisHostname, util.ProcessToken.Hash,
		)
		if err != nil {
			return false, log.Errore(err)
		}
		rows, err := sqlResult.RowsAffected()
		if err != nil {
			return false, log.Errore(err)
		}
		if rows > 0 {
			// Reaffirmed our own leadership
			return true, nil
		}
	}
	return false, nil
}

// GrabElection forcibly grabs leadership. Use with care!!
func GrabElection() error {
	if orcraft.IsRaftEnabled() {
		return log.Errorf("Cannot GrabElection on raft setup")
	}
	_, err := db.ExecOrchestrator(`
			replace into active_node (
					anchor, hostname, token, first_seen_active, last_seen_active
				) values (
					1, ?, ?, now(), now()
				)
			`,
		ThisHostname, util.ProcessToken.Hash,
	)
	return log.Errore(err)
}

// Reelect clears the way for re-elections. Active node is immediately demoted.
func Reelect() error {
	if orcraft.IsRaftEnabled() {
		orcraft.StepDown()
	}
	_, err := db.ExecOrchestrator(`delete from active_node where anchor = 1`)
	return log.Errore(err)
}

// ElectedNode returns the details of the elected node, as well as answering the question "is this process the elected one"?
func ElectedNode() (node NodeHealth, isElected bool, err error) {
	query := `
		select
			hostname,
			token,
			first_seen_active,
			last_seen_Active
		from
			active_node
		where
			anchor = 1
		`
	err = db.QueryOrchestratorRowsMap(query, func(m sqlutils.RowMap) error {
		node.Hostname = m.GetString("hostname")
		node.Token = m.GetString("token")
		node.FirstSeenActive = m.GetString("first_seen_active")
		node.LastSeenActive = m.GetString("last_seen_active")

		return nil
	})

	isElected = (node.Hostname == ThisHostname && node.Token == util.ProcessToken.Hash)
	return node, isElected, log.Errore(err)
}
