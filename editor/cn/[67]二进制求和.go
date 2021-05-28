//给你两个二进制字符串，返回它们的和（用二进制表示）。 
//
// 输入为 非空 字符串且只包含数字 1 和 0。 
//
// 
//
// 示例 1: 
//
// 输入: a = "11", b = "1"
//输出: "100" 
//
// 示例 2: 
//
// 输入: a = "1010", b = "1011"
//输出: "10101" 
//
// 
//
// 提示： 
//
// 
// 每个字符串仅由字符 '0' 或 '1' 组成。 
// 1 <= a.length, b.length <= 10^4 
// 字符串如果不是 "0" ，就都不含前导零。 
// 
// Related Topics 数学 字符串 
// 👍 614 👎 0

package cn

import (
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func addBinary(a string, b string) string {
	listA := strings.Split(a, "")
	listB := strings.Split(b, "")
	lenA := len(listA)
	lenB := len(listB)
	cluA, cluB := lenA-1, lenB-1
	var result []string
	carryBit, bitA, bitB := 0, 0, 0
	var err error
	for {
		if cluA < 0 && cluB < 0 {
			break
		}
		if cluA < 0 {
			bitA = 0
		} else {
			bitA, err = strconv.Atoi(listA[cluA])
			if err != nil {
				return ""
			}
		}
		if cluB < 0 {
			bitB = 0
		} else {
			bitB, err = strconv.Atoi(listB[cluB])
			if err != nil {
				return ""
			}
		}

		switch bitA + bitB + carryBit {
		case 3:
			result = append([]string{"1"}, result...)
			carryBit = 1
		case 2:
			result = append([]string{"0"}, result...)
			carryBit = 1
		case 1:
			result = append([]string{"1"}, result...)
			carryBit = 0
		case 0:
			result = append([]string{"0"}, result...)
			carryBit = 0
		default:
			return ""
		}
		cluA, cluB = cluA-1, cluB-1
	}
	if carryBit == 1{
		result = append([]string{"1"}, result...)
	}
	return strings.Join(result, "")
}

//leetcode submit region end(Prohibit modification and deletion)
