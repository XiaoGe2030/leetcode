package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func numWays(n int) int {
	if n <= 1 {
		return 1
	}

	a, b := 0, 1
	for n > 0 {
		tmp := a + b
		a = b
		b = tmp % 1000000007
		n--
	}
	return b
}
