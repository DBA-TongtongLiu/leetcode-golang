/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

//leetcode submit region begin(Prohibit modification and deletion)
func exchangeBits(num int) int {
	bina := fmt.Sprintf("%b", num)
	list := strings.Split(bina, "")
	length := len(list)
	if length%2 == 1 {
		list = append([]string{"0"}, list...)
		length ++
	}
	for i := 0; i <= length-2; i += 2 {
		list[i], list[i+1] = list[i+1], list[i]
	}
	newNum, _ := strconv.ParseInt(strings.Join(list, ""), 2, 32)

	return int(newNum)
}
