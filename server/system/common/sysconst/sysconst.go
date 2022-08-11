/*
 * @Date: 2022-03-24 08:36:18
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc: 系统固定常量
 * @LastEditTime: 2022-08-11 16:19:44
 * @FilePath: \server\system\common\sysconst\sysconst.go
 */
package sysconst

//系统表
var SysTables = []string{"sys_permission", "sys_role", "sys_role_permission", "sys_user", "sys_user_permission", "sys_user_role", "sys_config", "sys_sequence"}
var EmsConfigTables = []string{"ems_acq_client", "ems_acq_correct", "ems_acq_points", "ems_alarm_record", "ems_analysis_config", "ems_unit_price", "ems_units_config"}
var EmsDataTables = []string{"ems_acq_cost", "ems_acq_data"}

const CacheKey_ids string = "online"
