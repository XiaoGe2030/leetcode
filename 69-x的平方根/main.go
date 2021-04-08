package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vim-go")
}

//调用库函数
func mySqrt(x int) int {
	return int(math.Sqrt(float64(x)))
}

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	ans := int(math.Exp(0.5 * math.Log(float64(x))))
	if (ans+1)*(ans+1) <= x {
		return ans + 1
	}
	return ans
}
