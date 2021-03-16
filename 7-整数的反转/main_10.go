package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1534236469))
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
}

func reverse(x int) int {
	res := 0
	for x != 0 {
		//每次取末尾数字
		tmp := x % 10
		//判断是否 大于 最大32位整数
		if res > (math.MaxInt32/10) || (res == (math.MaxInt32/10) && tmp > 7) {
			return 0
		}
		//判断是否 小于 最小32位整数
		if res < (math.MinInt32/10) || (res == (math.MinInt32/10) && tmp < (-8)) {
			return 0
		}
		res = res*10 + tmp
		x /= 10
	}
	return res
}
