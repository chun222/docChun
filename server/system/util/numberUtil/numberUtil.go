/*
 * @Date: 2022-06-14 14:52:49
 * @LastEditors: 春贰
 * @gitee: https://gitee.com/chun22222222
 * @github: https://github.com/chun222
 * @Desc:
 * @LastEditTime: 2022-06-21 12:18:07
 * @FilePath: \server\system\util\numberUtil\numberUtil.go
 */
package numberUtil

import (
	"math"
	"strconv"
)

//保留n位小数
func RoundFloat(value float64, n int) float64 {
	value2, _ := strconv.ParseFloat(strconv.FormatFloat(math.Trunc(value*math.Pow10(n)+0.5)*math.Pow10(-n), 'f', n, 64), 64)
	return value2
}
