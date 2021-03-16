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
	var res, last int32
	for x != 0 {
		//每次取末尾数字
		tmp := x % 10
		last = res
		res = int32(int(res)*10 + tmp)
		//判断整数溢出
		if last != res/10 {
			return 0
		}
		x /= 10
	}
	return int(res)
}
