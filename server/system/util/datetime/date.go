/*
 * @Date: 2022-06-07 14:59:26
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-06-24 17:06:28
 * @FilePath: \server\system\util\datetime\date.go
 */
package datetime

import "time"

func Year(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.Year()
}

func Month(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return int(tmp.Month())
}

func Day(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.Day()
}

func YearDay(t ...time.Time) int {
	tmp := time.Now()
	if len(t) > 0 {
		tmp = t[0]
	}
	return tmp.YearDay()
}

// getYearMonthToDay 查询指定年份指定月份有多少天
// @params year int 指定年份
// @params month int 指定月份
func GetYearMonthToDay(year int, month int) int {
	// 有31天的月份
	day31 := map[int]string{
		1:  "",
		3:  "",
		5:  "",
		7:  "",
		8:  "",
		10: "",
		12: "",
	}
	if _, ok := day31[month]; ok {
		return 31
	}
	// 有30天的月份
	day30 := map[int]string{
		4:  "",
		6:  "",
		9:  "",
		11: "",
	}
	if _, ok := day30[month]; ok {
		return 30
	}
	// 计算是平年还是闰年
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		// 得出2月的天数
		return 29
	}
	// 得出2月的天数
	return 28
}
