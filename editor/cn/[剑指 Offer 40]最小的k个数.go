//输入整数数组 arr ，找出其中最小的 k 个数。
//例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4。
//
// 
//
// 示例 1： 
//
// 输入：arr = [3,2,1], k = 2
//输出：[1,2] 或者 [2,1]
// 
//
// 示例 2： 
//
// 输入：arr = [0,1,2,1], k = 1
//输出：[0] 
//
// 
//
// 限制： 
//
// 
// 0 <= k <= arr.length <= 10000 
// 0 <= arr[i] <= 10000 
// 
// Related Topics 堆 分治算法 
// 👍 251 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func getLeastNumbers(arr []int, k int) []int {
	mm := make(map[int]int)
	for _, a := range arr {
		if v, ok := mm[a]; ok {
			mm[a] = v + 1
		} else {
			mm[a] = 1
		}
	}
	min := 9999999
	cur := 0
	for {
		if cur >= k{
			break
		}
		minK,minV := getMin(mm)

	}
	return nil
}

func getMin(m map[int]int) (key, value int) {
	min := 9999999999999
	for k, _ := range m {
		if k < min {
			min = k
		}
	}
	return min, m[min]
}

//leetcode submit region end(Prohibit modification and deletion)
