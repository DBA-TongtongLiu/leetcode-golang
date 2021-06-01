/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:30 下午
**/

package leetcode

//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	length := len(flowerbed)
	for i, v := range flowerbed {
		if v == 1 {
			continue
		}

		if i == 0 {
			if i+1 < length && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		} else if i == length-1 && flowerbed[i-1] == 0 {
			flowerbed[i] = 1
			n--
		} else {
			if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n--
			}
		}

	}
	if n > 0 {
		return false
	}
	return true
}