//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。 
//
// 有效字符串需满足： 
//
// 
// 左括号必须用相同类型的右括号闭合。 
// 左括号必须以正确的顺序闭合。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：s = "()"
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：s = "()[]{}"
//输出：true
// 
//
// 示例 3： 
//
// 
//输入：s = "(]"
//输出：false
// 
//
// 示例 4： 
//
// 
//输入：s = "([)]"
//输出：false
// 
//
// 示例 5： 
//
// 
//输入：s = "{[]}"
//输出：true 
//
// 
//
// 提示： 
//
// 
// 1 <= s.length <= 104 
// s 仅由括号 '()[]{}' 组成 
// 
// Related Topics 栈 字符串 
// 👍 2405 👎 0

package cn

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	var stack []string
	sArray := strings.Split(s, "")
	for _, s := range sArray {
		switch s {
		case "(", "[", "{":
			stack = append(stack, s)
		case ")":
			if stack == nil || len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "(" {
				stack = stack[0 : len(stack)-1]
				continue
			} else {
				return false
			}
		case "]":
			if stack == nil || len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "[" {
				stack = stack[0 : len(stack)-1]
				continue
			} else {
				return false
			}
		case "}":
			if stack == nil || len(stack) == 0 {
				return false
			}
			if stack[len(stack)-1] == "{" {
				stack = stack[0 : len(stack)-1]
				continue
			} else {
				return false
			}
		default:
			return false
		}
	}
	if len(stack) == 0 {
		return true
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
