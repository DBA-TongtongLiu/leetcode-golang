/**
* @Author: TongTongLiu
* @Date: 2021/5/17 2:38 下午
**/

package test

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	var negative bool
	if x < 0 {
		negative = true
		x = -x
	}
	str := strconv.Itoa(x)
	r := []rune(str)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	xRev, _ := strconv.Atoi(string(r))
	if negative && float64(xRev) > math.Pow(2, 31)-1 {
		return 0
	}
	if !negative && float64(xRev) > math.Pow(2, 31) {
		return 0
	}
	if negative {
		return -xRev
	}
	return xRev
}
