/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

import "strings"

//leetcode submit region begin(Prohibit modification and deletion)
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var strList []*string
	column, flag := 0, 1
	for i := 0; i < numRows; i++ {
		str := ""
		strList = append(strList, &str)
	}
	for _, s := range strings.Split(s, "") {
		*(strList[column]) += s
		column += flag
		if column+1 == numRows {
			flag = -1
		} else if column == 0 {
			flag = 1
		}

	}
	var str string
	for _, s := range strList {
		str += *s
	}
	return str
}
