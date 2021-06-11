package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func jump1(nums []int) int {

	n := len(nums)

	pos = n - 1
	var step, end int
	for pos > 0 {
		for i := 0; i <= pos; i++ {

			if nums[i]+i >= pos {
				step++
				pos = i
				break
			}
		}
	}
	return step

}

func jump(nums []int) int {

	n := len(nums)
	pos := 0
	var step, end int
	for i := 0; i < n-1; i++ {
		pos = max(pos, nums[i])
		if i == end {
			end = pos
			step++
		}

	}

	return step
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
