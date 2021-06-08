//给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。 
//
// 在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。 
//
// 注意: 
//假设字符串的长度不会超过 1010。 
//
// 示例 1: 
//
// 
//输入:
//"abccccdd"
//
//输出:
//7
//
//解释:
//我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
// 
// Related Topics 哈希表 
// 👍 302 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func longestPalindrome(s string) int {
	m := make(map[int32]int)
	for _, i := range s {
		if v, ok := m[i]; ok {
			m[i] = v + 1
		} else {
			m[i] = 1
		}
	}
	var max int
	var flag bool
	for _, v := range m {
		if v%2 == 1 {
			flag = true
			v = v - 1
		}
		max += v
	}
	if flag {
		max++
	}
	return max
}

//leetcode submit region end(Prohibit modification and deletion)
