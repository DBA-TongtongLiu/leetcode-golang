/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func findPoisonedDuration(timeSeries []int, duration int) int {
	var continueTime, endTime int
	for _, t := range timeSeries {
		if t == 0 {
		} else if endTime >= t {
			continueTime -= endTime - t + 1
		}
		endTime = t + duration - 1
		continueTime += duration
	}
	return continueTime
}
