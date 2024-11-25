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

/*
 * 审计日志至少包涵 3 个维度
 * 1、实例 InstanceKey
 * 2、时间 AuditTimestamp
 * 3、类型 & 信息   AuditType & Message
 */
// Audit presents a single audit entry (namely in the database)
type Audit struct {
	AuditId          int64
	AuditTimestamp   string
	AuditType        string
	AuditInstanceKey InstanceKey
	Message          string
}
