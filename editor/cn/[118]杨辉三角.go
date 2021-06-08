//给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。 
//
// 
//
// 在杨辉三角中，每个数是它左上方和右上方的数的和。 
//
// 示例: 
//
// 输入: 5
//输出:
//[
//     [1],
//    [1,1],
//   [1,2,1],
//  [1,3,3,1],
// [1,4,6,4,1]
//] 
// Related Topics 数组 
// 👍 501 👎 0

package cn

//leetcode submit region begin(Prohibit modification and deletion)
func generate(numRows int) [][]int {
	var result [][]int
	var last []int
	for i := 1; i <= numRows; i++ {
		switch i {
		case 1:
			last = []int{1}
		case 2:
			last = []int{1, 1}
		default:
			last = getNext(last)
		}
		result = append(result, last)
	}
	return result
}

func getNext(in []int) (out []int) {

	for i, _ := range in {
		if i == 0 {
			out = append(out, 1)
		} else {
			out = append(out, in[i-1]+in[i])
		}
	}
	return append(out, 1)
}

//leetcode submit region end(Prohibit modification and deletion)
