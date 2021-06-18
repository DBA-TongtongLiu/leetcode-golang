//给定一个字符串来代表一个学生的出勤记录，这个记录仅包含以下三个字符： 
//
// 
// 'A' : Absent，缺勤 
// 'L' : Late，迟到 
// 'P' : Present，到场 
// 
//
// 如果一个学生的出勤记录中不超过一个'A'(缺勤)并且不超过两个连续的'L'(迟到),那么这个学生会被奖赏。 
//
// 你需要根据这个学生的出勤记录判断他是否会被奖赏。 
//
// 示例 1: 
//
// 输入: "PPALLP"
//输出: True
// 
//
// 示例 2: 
//
// 输入: "PPALLL"
//输出: False
// 
// Related Topics 字符串 
// 👍 73 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func checkRecord(s string) bool {
	var numA, consL int
	for _, ss := range []byte(s) {
		if string(ss) == "A" {
			numA++
			if numA > 1 {
				return false
			}
			consL = 0
			continue
		}
		if string(ss) == "L" {
			consL++
			if consL > 2 {
				return false
			}
			continue
		}
		consL = 0
	}
	return true
}

//leetcode submit region end(Prohibit modification and deletion)
