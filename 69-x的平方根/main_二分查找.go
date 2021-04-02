package main

import (
	"fmt"
)

func main() {
	fmt.Println(mySqrt(1))
}

//调用库函数
func mySqrt(x int) int {
	l, r := 0, x
	var ans int
	for l <= r {
		mid := (r + l) / 2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
