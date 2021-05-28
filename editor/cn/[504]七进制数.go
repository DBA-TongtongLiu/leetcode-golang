//给定一个整数，将其转化为7进制，并以字符串形式输出。 
//
// 示例 1: 
//
// 
//输入: 100
//输出: "202"
// 
//
// 示例 2: 
//
// 
//输入: -7
//输出: "-10"
// 
//
// 注意: 输入范围是 [-1e7, 1e7] 。 
// 👍 85 👎 0

package cn

import (
	"math"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	var positive bool
	if num >= 0 {
		positive = true
	} else {
		num = int(math.Abs(float64(num)))
	}
	value, sevenValue := num, 0
	for num > 0 {
		count := 0
		for value > 6 {
			value = value / 7
			count++
		}
		sevenValue = sevenValue + value*int(math.Pow(10, float64(count)))
		num = num - value*int(math.Pow(7, float64(count)))
		value = num
	}
	if !positive {
		return "-" + strconv.Itoa(sevenValue)
	}
	return strconv.Itoa(sevenValue)
}

//leetcode submit region end(Prohibit modification and deletion)
