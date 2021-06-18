//数组中占比超过一半的元素称之为主要元素。给定一个整数数组，找到它的主要元素。若没有，返回-1。 
//
// 示例 1： 
//
// 输入：[1,2,5,9,5,9,5,5,5]
//输出：5 
//
// 
//
// 示例 2： 
//
// 输入：[3,2]
//输出：-1 
//
// 
//
// 示例 3： 
//
// 输入：[2,2,1,1,1,2,2]
//输出：2 
//
// 
//
// 说明： 
//你有办法在时间复杂度为 O(N)，空间复杂度为 O(1) 内完成吗？ 
// Related Topics 位运算 数组 分治算法 
// 👍 85 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	m := make(map[int]int)
	for _, n := range nums {
		if v, ok := m[n]; ok {
			m[n] = v + 1
			if m[n] > length/2.0 {
				return n
			}
		} else {
			m[n] = 1
		}

	}
	return -1
}

//leetcode submit region end(Prohibit modification and deletion)
