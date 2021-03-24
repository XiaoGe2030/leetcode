package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func findKthLargest(nums []int, k int) int {
	for i := 1; i <= k; i++ {
		for j := 0; j < len(nums)-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	return nums[len(nums)-k]
}
