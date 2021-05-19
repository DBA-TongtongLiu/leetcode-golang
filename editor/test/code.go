/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func maximumPopulation(logs [][]int) int {
	// 现在 2021 创建长度为 2021-1950+1 = 72 的数组
	var pop [999]int
	var sub int
	for _, l := range logs {
		sub = l[0] - 1950
		pop[sub] = pop[sub] + 1
		sub = l[1] - 1950
		pop[sub] = pop[sub] - 1
	}
	var sum, max int
	for i, v := range pop {
		if i == 0 {
			pop[i] = v
			sum = v
			max = v
			sub = i
		} else {
			sum += v
			pop[i] = sum
			if sum > max {
				max = sum
				sub = i
			}
		}
	}
	return 1950 + sub
}
