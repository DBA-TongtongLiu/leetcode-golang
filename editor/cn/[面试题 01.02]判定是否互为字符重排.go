//给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。 
//
// 示例 1： 
//
// 输入: s1 = "abc", s2 = "bca"
//输出: true 
// 
//
// 示例 2： 
//
// 输入: s1 = "abc", s2 = "bad"
//输出: false
// 
//
// 说明： 
//
// 
// 0 <= len(s1) <= 100 
// 0 <= len(s2) <= 100 
// 
// Related Topics 数组 字符串 
// 👍 33 👎 0
package cn

import (
	"reflect"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func CheckPermutation(s1 string, s2 string) bool {
	strL1 := strings.Split(s1, "")
	strL2 := strings.Split(s2, "")
	strM1 := make(map[string]int)
	strM2 := make(map[string]int)
	for _, s := range strL1 {
		if sum, ok := strM1[s]; !ok {
			strM1[s] = 1
			continue
		} else {
			strM1[s] = sum + 1
		}
	}
	// 练手项目，不抽象方法了
	for _, s := range strL2 {
		if sum, ok := strM2[s]; !ok {
			strM2[s] = 1
			continue
		} else {
			strM2[s] = sum + 1
		}
	}
	return reflect.DeepEqual(strM1, strM2)
}

//leetcode submit region end(Prohibit modification and deletion)
