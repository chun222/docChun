/*
 * @Date: 2022-02-14 15:05:35
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-04-11 15:33:19
 * @FilePath: \server\system\model\DbModel\sysTools.go
 */
package DbModel

type SysSequence struct {
	Name  string `json:"name" gorm:"primaryKey;not null;comment:name"`
	Value int64  `json:"value" gorm:"default:0;not null;comment:value"`
}
