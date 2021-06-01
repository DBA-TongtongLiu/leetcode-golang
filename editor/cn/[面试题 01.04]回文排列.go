//给定一个字符串，编写一个函数判定其是否为某个回文串的排列之一。 
//
// 回文串是指正反两个方向都一样的单词或短语。排列是指字母的重新排列。 
//
// 回文串不一定是字典当中的单词。 
//
// 
//
// 示例1： 
//
// 输入："tactcoa"
//输出：true（排列有"tacocat"、"atcocta"，等等）
// 
//
// 
// Related Topics 哈希表 字符串 
// 👍 52 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func canPermutePalindrome(s string) bool {
	m := make(map[int32]bool)
	for _, ss := range s {
		if _, ok := m[ss]; ok {
			delete(m, ss)
		} else {
			m[ss] = true
		}
	}
	if len(m) > 1 {
		return false
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
