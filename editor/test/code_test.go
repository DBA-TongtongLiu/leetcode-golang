/**
* @Author: TongTongLiu
* @Date: 2021/5/17 3:27 下午
**/

package leetcode

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	logs := [][]int{{1982, 1998}, {2013, 2042}, {2010, 2035}, {2022, 2050}, {2047, 2048}}
	fmt.Println(maximumPopulation(logs))
}
