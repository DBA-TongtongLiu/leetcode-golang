//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。 
//
// 如果反转后整数超过 32 位的有符号整数的范围 [−231, 231 − 1] ，就返回 0。 
//假设环境不允许存储 64 位整数（有符号或无符号）。
//
// 
//
// 示例 1： 
//
// 
//输入：x = 123
//输出：321
// 
//
// 示例 2： 
//
// 
//输入：x = -123
//输出：-321
// 
//
// 示例 3： 
//
// 
//输入：x = 120
//输出：21
// 
//
// 示例 4： 
//
// 
//输入：x = 0
//输出：0
// 
//
// 
//
// 提示： 
//
// 
// -231 <= x <= 231 - 1 
// 
// Related Topics 数学 
// 👍 2807 👎 0

package cn

import (
	"math"
	"strconv"
)

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x = -x
	}
	str := strconv.Itoa(x)
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	xRev, _ := strconv.Atoi(string(r))
	if negative && float64(xRev) > math.Pow(2, 31)-1 {
		return 0
	}
	if !negative && float64(xRev) > math.Pow(2, 31) {
		return 0
	}
	if negative {
		return -xRev
	}
	return xRev
}

//leetcode submit region end(Prohibit modification and deletion)
