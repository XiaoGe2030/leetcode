package main

import (
	"fmt"
)

func main() {
	fmt.Println(waysToStep(61))
}

/*
三步问题。有个小孩正在上楼梯，楼梯有n阶台阶，小孩一次可以上1阶、2阶或3阶。实现一种方法，计算小孩有多少种上楼梯的方式。结果可能很大，你需要对结果模1000000007。
*/
func waysToStep(n int) int {
	a, b, c := 0, 0, 1
	for n > 0 {
		tmp := a + b + c
		a = b
		b = c
		c = tmp % 1000000007
		n--
	}
	return c
}
