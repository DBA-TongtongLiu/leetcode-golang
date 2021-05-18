//给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。 
//
// 你可以假设数组中无重复元素。 
//
// 示例 1: 
//
// 输入: [1,3,5,6], 5
//输出: 2
// 
//
// 示例 2: 
//
// 输入: [1,3,5,6], 2
//输出: 1
// 
//
// 示例 3: 
//
// 输入: [1,3,5,6], 7
//输出: 4
// 
//
// 示例 4: 
//
// 输入: [1,3,5,6], 0
//输出: 0
// 
// Related Topics 数组 二分查找 
// 👍 912 👎 0

package cn

import "math"

//leetcode submit region begin(Prohibit modification and deletion)
func searchInsert(nums []int, target int) int {
	length := len(nums)
	begin, end, mid := 0, length-1, 0
	if nums[begin] > target {
		return 0
	}
	if nums[end] < target {
		return end + 1
	}
	for begin < end {
		mid = int(math.Floor(float64(begin+end+1) / 2))
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			begin = mid
		} else {
			end = mid
		}
		if begin+1 == end {
			if nums[begin] == target {
				return begin
			} else if nums[begin] > target {
				return begin
			} else {
				return begin + 1
			}
		}
	}
	return 0
}

//leetcode submit region end(Prohibit modification and deletion)
