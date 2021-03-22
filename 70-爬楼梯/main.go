package main

import "fmt"

func main() {
<<<<<<< HEAD
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {
	var a, b int
	var c int
	c = 1
=======
	fmt.Println(climbStairs(4))
}

func climbStairs(n int) int {
	a, b, c := 0, 0, 1
>>>>>>> 20400e72f6b5c14362dc03729df340c899d1cd9f
	for i := 1; i <= n; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}
