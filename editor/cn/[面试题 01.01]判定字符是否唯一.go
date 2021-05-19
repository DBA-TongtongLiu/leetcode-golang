//实现一个算法，确定一个字符串 s 的所有字符是否全都不同。 
//
// 示例 1： 
//
// 输入: s = "leetcode"
//输出: false 
// 
//
// 示例 2： 
//
// 输入: s = "abc"
//输出: true
// 
//
// 限制： 
// 
// 0 <= len(s) <= 100 
// 如果你不使用额外的数据结构，会很加分。 
// 
// Related Topics 数组 
// 👍 110 👎 0

package cn

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func isUnique(str string) bool {
	m := make(map[string]bool)
	for _, s:= range strings.Split(str, ""){
		if _, ok := m[s];ok{
			return false
		}else {
			m[s] = true
		}
	}
	return true
}
//leetcode submit region end(Prohibit modification and deletion)
