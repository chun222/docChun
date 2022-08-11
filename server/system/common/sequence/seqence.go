/*
 * @Date: 2022-04-11 16:51:26
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:获取唯一序列
 * @LastEditTime: 2022-04-12 09:03:37
 * @FilePath: \server\system\common\sequence\seqence.go
 */
package sequence

import (
	"strings"

	"chunDoc/system/core/db"
	"chunDoc/system/model/DbModel"
	"chunDoc/system/util/convert"
	"chunDoc/system/util/str"
)

func Get(name string) (error, int64) {
	var nowSEQ int64 = 0
	tx := db.Instance().Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err, 0
	}

	var seq DbModel.SysSequence

	//开启mysql锁
	tx.Exec("set autocommit=0")

	if err := tx.Set("gorm:query_option", "FOR UPDATE").Limit(1).Where("name = ?", name).Find(&seq).Error; err != nil {
		tx.Rollback()
		return err, 0
	}

	if seq.Name != "" {
		//更新
		seq.Value++
		tx.Save(&seq)
		nowSEQ = seq.Value
	} else {
		// 新增
		_seq := DbModel.SysSequence{Name: name, Value: 1}
		tx.Create(&_seq)
		nowSEQ = 1

	}

	// commit事务，释放锁
	if err := tx.Commit().Error; err != nil {
		return err, 0
	}

	return nil, nowSEQ
}

//将序列组合成指定流水号
func GetString(pre string, length int) (error, string) {
	if err, seqnum := Get(strings.ToUpper(pre)); err != nil {
		return err, ""
	} else {
		return nil, strings.ToUpper(pre) + str.Leftpad(convert.String(seqnum), "0", length-len(pre))
	}

}
