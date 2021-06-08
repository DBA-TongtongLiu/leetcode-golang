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
func isMonotonic(nums []int) bool {
	if len(nums) <= 2 {
		return true
	}
	// flag == true 单调递增
	// flag == false 单调递减
	flag := nums[0] <= nums[1]

	for i, _ := range nums {
		if i == 1 || i == 0 {
			continue
		}
		if nums[i] == nums[i-1]{
			continue
		}
		if flag != (nums[i] > nums[i-1]) {
			return false
		}
	}
	return true
}
