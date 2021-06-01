//给定一个字符串，判断该字符串中是否可以通过重新排列组合，形成一个回文字符串。 
//
// 示例 1： 
//
// 输入: "code"
//输出: false 
//
// 示例 2： 
//
// 输入: "aab"
//输出: true 
//
// 示例 3： 
//
// 输入: "carerac"
//输出: true 
// Related Topics 哈希表 
// 👍 39 👎 0
package cn

//leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	m := make(map[int32]int)
	for _, j := range s {
		if v, ok := m[j]; ok {
			m[j] = v + 1
		} else {
			m[j] = 1
		}
	}
	var cardinal int
	for _, v := range m {
		if v%2 == 1 {
			cardinal++
		}
		if cardinal > 1 {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
