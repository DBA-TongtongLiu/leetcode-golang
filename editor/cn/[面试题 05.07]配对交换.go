//配对交换。编写程序，交换某个整数的奇数位和偶数位，尽量使用较少的指令（也就是说，位0与位1交换，位2与位3交换，以此类推）。 
//
// 示例1: 
//
// 
// 输入：num = 2（或者0b10）
// 输出 1 (或者 0b01)
// 
//
// 示例2: 
//
// 
// 输入：num = 3
// 输出：3
// 
//
// 提示: 
//
// 
// num的范围在[0, 2^30 - 1]之间，不会发生整数溢出。 
// 
// Related Topics 位运算 
// 👍 41 👎 0

package cn

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
		length++
	}
	for i := 0; i <= length-2; i += 2 {
		list[i], list[i+1] = list[i+1], list[i]
	}
	newNum, _ := strconv.ParseInt(strings.Join(list, ""), 2, 32)

	return int(newNum)
}

//leetcode submit region end(Prohibit modification and deletion)
