/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func missingNumber(nums []int) int {
	var sum, max int
	for _, i := range nums {
		sum += i
		if i > max {
			max = i
		}
	}
	sup := (max + 1) * max / 2
	if sup != sum {
		return sup - sum
	}
	length := len(nums)
	if max == length {
		return 0
	} else {
		return max + 1
	}

	return - sum
}