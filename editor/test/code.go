/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

import (
	"math"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func convertToBase7(num int) string {
	var positive bool
	if num > 0 {
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
