package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
		}
	}

	return res
}
