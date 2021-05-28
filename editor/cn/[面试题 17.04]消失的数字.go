//数组nums包含从0到n的所有整数，但其中缺了一个。请编写代码找出那个缺失的整数。你有办法在O(n)时间内完成吗？ 
//
// 注意：本题相对书上原题稍作改动 
//
// 示例 1： 
//
// 输入：[3,0,1]
//输出：2 
//
// 
//
// 示例 2： 
//
// 输入：[9,6,4,2,3,5,7,0,1]
//输出：8
// 
// Related Topics 位运算 数组 数学 
// 👍 40 👎 0
package cn

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

//leetcode submit region end(Prohibit modification and deletion)
// 这个方法已经达到 双 100%, 但我还是想要看看抑或怎么写
//func missingNumber(nums []int) int {
//	var sum, max int
//
//	for _, i := range nums {
//		sum += i
//		if i > max {
//			max = i
//		}
//	}
//
//	sup := (max + 1) * max / 2
//	if sup != sum {
//		return sup - sum
//	}
//	length := len(nums)
//	if max == length {
//		return 0
//	} else {
//		return max + 1
//	}
//
//	return - sum
//}