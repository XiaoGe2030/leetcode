package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	max := 0
	n := len(height)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			capw := capWater(height[i], height[j], i, j)
			if capw > max {
				max = capw
			}
		}
	}
	return max
}

func capWater(m, n, x, y int) int {
	a, b := m, y-x
	if m > n {
		a = n
	}
	return a * b
}
