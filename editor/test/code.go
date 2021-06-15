/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

//leetcode submit region begin(Prohibit modification and deletion)
func isPalindrome(s string) bool {
	var strNew string
	for _, l := range []byte(s) {
		if string(l) >= "A" && string(l) <= "Z" {
			strNew += string(l+32)
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
