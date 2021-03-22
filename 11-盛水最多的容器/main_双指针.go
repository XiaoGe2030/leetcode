package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 1}))
}

func maxArea(height []int) int {
	n := len(height)
	left, right := 0, n-1
	max := capW(height[left], height[right], left, right)
	for left < right {
		// left 比前一个高才需要计算
		if left > 0 && height[left] > height[left-1] {
			capW := capW(height[left], height[right], left, right)
			if max < capW {
				max = capW
			}
		}
		// right 比前一个高才需要计算
		if right < n-1 && height[right] > height[right+1] {
			capW := capW(height[left], height[right], left, right)
			if max < capW {
				max = capW
			}
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return max
}

func capW(n, m, x, y int) int {
	a, b := n, y-x
	if m < n {
		a = m
	}

	return a * b
}
