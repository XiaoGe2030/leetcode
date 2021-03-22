package main

import (
	"fmt"
)

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	ans := 0
	for l < r {
		area := capW(height[l], height[r], l, r)
		if ans < area {
			ans = area
		}
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return ans
}

func capW(l, r, x, y int) int {
	a, b := l, y-x
	if l > r {
		a = r
	}

	return a * b
}
