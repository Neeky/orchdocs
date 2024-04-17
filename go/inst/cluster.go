/*
   Copyright 2014 Outbrain Inc.

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

package inst

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/openark/orchestrator/go/config"
	"github.com/openark/orchestrator/go/kv"
	"github.com/outbrain/golib/log"
)

func GetClusterMasterKVKey(clusterAlias string) string {
	return fmt.Sprintf("%s%s", config.Config.KVClusterMasterPrefix, clusterAlias)
}

func getClusterMasterKVPair(clusterAlias string, masterKey *InstanceKey) *kv.KVPair {
	if clusterAlias == "" {
		return nil
	}
	if masterKey == nil {
		return nil
	}
	return kv.NewKVPair(GetClusterMasterKVKey(clusterAlias), masterKey.StringCode())
}

// GetClusterMasterKVPairs returns all KV pairs associated with a master. This includes the
// full identity of the master as well as a breakdown by hostname, port, ipv4, ipv6
func GetClusterMasterKVPairs(clusterAlias string, masterKey *InstanceKey) (kvPairs [](*kv.KVPair)) {
	masterKVPair := getClusterMasterKVPair(clusterAlias, masterKey)
	if masterKVPair == nil {
		return kvPairs
	}
	kvPairs = append(kvPairs, masterKVPair)

	addPair := func(keySuffix, value string) {
		key := fmt.Sprintf("%s/%s", masterKVPair.Key, keySuffix)
		kvPairs = append(kvPairs, kv.NewKVPair(key, value))
	}

	addPair("hostname", masterKey.Hostname)
	addPair("port", fmt.Sprintf("%d", masterKey.Port))
	if ipv4, ipv6, err := readHostnameIPs(masterKey.Hostname); err == nil {
		addPair("ipv4", ipv4)
		addPair("ipv6", ipv6)
	}
	return kvPairs
}

/*
 * 根据 config.Config.ClusterNameToAlias 中配置的转换规则，把 clusterName 转换成对应的 clusterAlias
 * 如果没有任何一条规则与之匹配就返回  "" 串
 */
// mappedClusterNameToAlias attempts to match a cluster with an alias based on
// configured ClusterNameToAlias map
func mappedClusterNameToAlias(clusterName string) string {
	for pattern, alias := range config.Config.ClusterNameToAlias {
		if pattern == "" {
			// sanity
			continue
		}
		if matched, _ := regexp.MatchString(pattern, clusterName); matched {
			return alias
		}
	}
	return ""
}

/*
 * 集群信息对象
 *
 * 从它的设计来看一个集群应该有 3 个维度可以标识他了
 * 1、clusterName   集群名
 * 2、clusterAlias  对人类更加友好的一个名字
 * 3、clusterDomain 集群 master 结点的 域名、vip、dns-A-记录
 */
// ClusterInfo makes for a cluster status/info summary
type ClusterInfo struct {
	ClusterName                            string
	ClusterAlias                           string // Human friendly alias
	ClusterDomain                          string // CNAME/VIP/A-record/whatever of the master of this cluster
	CountInstances                         uint   // 实例数
	HeuristicLag                           int64  // 探测间隔
	HasAutomatedMasterRecovery             bool   // 是否命中 config.Config.RecoverMasterClusterFilters 中配置的正则
	HasAutomatedIntermediateMasterRecovery bool   // 是否命中 config.Config.RecoverIntermediateMasterClusterFilters 中配置的正则
}

/*
 * 从配置文件中读取 config.Config.RecoverMasterClusterFilters 和 config.Config.RecoverIntermediateMasterClusterFilters 这个两项的值
 * 1、只有能被  config.Config.RecoverMasterClusterFilters 中给定的 pattern 匹配的实例才会进行宕机切换
 * 2、config.Config.RecoverIntermediateMasterClusterFilters 同理
 *
 *
 */
// ReadRecoveryInfo
func (this *ClusterInfo) ReadRecoveryInfo() {
	log.Warning("enter ClusterInfo.ReadRecoveryInfo", "HasAutomatedMasterRecovery", this.HasAutomatedMasterRecovery)
	log.Warning("enter ClusterInfo.ReadRecoveryInfo", "HasAutomatedIntermediateMasterRecovery", this.HasAutomatedIntermediateMasterRecovery)

	this.HasAutomatedMasterRecovery = this.filtersMatchCluster(config.Config.RecoverMasterClusterFilters)
	this.HasAutomatedIntermediateMasterRecovery = this.filtersMatchCluster(config.Config.RecoverIntermediateMasterClusterFilters)

	log.Warning("exit ClusterInfo.ReadRecoveryInfo", "HasAutomatedMasterRecovery", this.HasAutomatedMasterRecovery)
	log.Warning("exit ClusterInfo.ReadRecoveryInfo", "HasAutomatedIntermediateMasterRecovery", this.HasAutomatedIntermediateMasterRecovery)
}

// filtersMatchCluster will see whether the given filters match the given cluster details
func (this *ClusterInfo) filtersMatchCluster(filters []string) bool {
	log.Warning("enter ClusterInfo.filtersMatchCluster:", "filters", filters)
	log.Warning("enter ClusterInfo.filtersMatchCluster:", "this.ClusterName", this.ClusterName, "this.ClusterAlias", this.ClusterAlias)
	for _, filter := range filters {
		if filter == this.ClusterName {
			return true
		}
		if filter == this.ClusterAlias {
			return true
		}
		if strings.HasPrefix(filter, "alias=") {
			// Match by exact cluster alias name
			alias := strings.SplitN(filter, "=", 2)[1]
			if alias == this.ClusterAlias {
				return true
			}
		} else if strings.HasPrefix(filter, "alias~=") {
			// Match by cluster alias regex
			aliasPattern := strings.SplitN(filter, "~=", 2)[1]
			if matched, _ := regexp.MatchString(aliasPattern, this.ClusterAlias); matched {
				return true
			}
		} else if filter == "*" {
			return true
		} else if matched, _ := regexp.MatchString(filter, this.ClusterName); matched && filter != "" {
			return true
		}
	}
	log.Warning("exit ClusterInfo.filtersMatchCluster:", "return", strconv.FormatBool(false))
	return false
}

// ApplyClusterAlias updates the given clusterInfo's ClusterAlias property
func (this *ClusterInfo) ApplyClusterAlias() {
	if this.ClusterAlias != "" && this.ClusterAlias != this.ClusterName {
		// Already has an alias; abort
		return
	}
	if alias := mappedClusterNameToAlias(this.ClusterName); alias != "" {
		this.ClusterAlias = alias
	}
}
