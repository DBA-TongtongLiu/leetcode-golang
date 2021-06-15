//给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。 
//
// 说明：本题中，我们将空字符串定义为有效的回文串。 
//
// 示例 1: 
//
// 输入: "A man, a plan, a canal: Panama"
//输出: true
// 
//
// 示例 2: 
//
// 输入: "race a car"
//输出: false
// 
// Related Topics 双指针 字符串 
// 👍 387 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	var strNew string
	for _, l := range []byte(s) {
		if string(l) >= "A" && string(l) <= "Z" {
			strNew += string(l + 32)
			continue
		}
		if string(l) >= "a" && string(l) <= "z" {
			strNew += string(l)
			continue
		}
	}
	list := []byte(strNew)
	length := len(list)
	for i, l := range list {
		if i > length/2 {
			break
		}
		if l != list[length-1-i] {
			return false
		}
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
