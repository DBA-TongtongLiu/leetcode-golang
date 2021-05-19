//给你一个二维整数数组 logs ，其中每个 logs[i] = [birthi, deathi] 表示第 i 个人的出生和死亡年份。 
//
// 年份 x 的 人口 定义为这一年期间活着的人的数目。第 i 个人被计入年份 x 的人口需要满足：x 在闭区间 [birthi, deathi - 1] 内
//。注意，人不应当计入他们死亡当年的人口中。 
//
// 返回 人口最多 且 最早 的年份。 
//
// 
//
// 示例 1： 
//
// 输入：logs = [[1993,1999],[2000,2010]]
//输出：1993
//解释：人口最多为 1 ，而 1993 是人口为 1 的最早年份。
// 
//
// 示例 2： 
//
// 输入：logs = [[1950,1961],[1960,1971],[1970,1981]]
//输出：1960
//解释： 
//人口最多为 2 ，分别出现在 1960 和 1970 。
//其中最早年份是 1960 。 
//
// 
//
// 提示： 
//
// 
// 1 <= logs.length <= 100 
// 1950 <= birthi < deathi <= 2050 
// 
// Related Topics 数组 
// 👍 11 👎 0
package cn

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

//leetcode submit region end(Prohibit modification and deletion)
// 比较暴力的方法
//func maximumPopulation(logs [][]int) int {
//	pop := make(map[int]int)
//	var max int
//	for _, l := range logs {
//		for y := l[0]; y < l[1]; y++ {
//			if v, ok := pop[y]; ok {
//				pop[y] = v + 1
//			} else {
//				pop[y] = 1
//			}
//			if pop[y] > max {
//				max = pop[y]
//			}
//		}
//	}
//	min := 9999
//	for k, v := range pop {
//		if v == max {
//			if k < min {
//				min = k
//			}
//		}
//	}
//
//	return min
//}
