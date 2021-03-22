package main

import "fmt"

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	var a, b int
	var c int
	c = 1
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}
