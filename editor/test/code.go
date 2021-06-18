/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement(nums []int) int {
	length := len(nums)
	m := make(map[int]int)
	for _, n := range nums {
		if v, ok := m[n]; ok {
			m[n] = v + 1
			if m[n] >length/2.0 {
				return n
			}
		} else {
			m[n] = 1
		}

	}
	return -1
}
