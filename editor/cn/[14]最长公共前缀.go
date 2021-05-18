//编写一个函数来查找字符串数组中的最长公共前缀。 
//
// 如果不存在公共前缀，返回空字符串 ""。 
//
// 
//
// 示例 1： 
//
// 
//输入：strs = ["flower","flow","flight"]
//输出："fl"
// 
//
// 示例 2： 
//
// 
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。 
//
// 
//
// 提示： 
//
// 
// 0 <= strs.length <= 200 
// 0 <= strs[i].length <= 200 
// strs[i] 仅由小写英文字母组成 
// 
// Related Topics 字符串 
// 👍 1603 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 && len([]rune(strs[0])) == 1{
		return strs[0]
	}
	var arrNew [][]rune
	minLen := 0
	for i, s := range strs {
		sRune := []rune(s)
		sLen := len(sRune)
		if i == 0 {
			minLen = sLen
		}
		if minLen > sLen {
			minLen = sLen
		}
		arrNew = append(arrNew, sRune)
	}
	var commRune []rune
	lenArr := len(arrNew)
	var endFlag bool
	for i := 0; i < minLen; i++ {
		var commR rune
		for j, r := range arrNew {
			if j == 0 {
				commR = r[i]
				continue
			}
			if commR != r[i] {
				endFlag = true
				break
			}
			if j == lenArr-1 {
				commRune = append(commRune, commR)
				break
			}
			continue
		}
		if endFlag{
			break
		}
	}
	return string(commRune)
}

//leetcode submit region end(Prohibit modification and deletion)
