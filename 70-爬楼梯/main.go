package main

import "fmt"

func main() {
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	a, b, c := 0, 0, 1
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}
