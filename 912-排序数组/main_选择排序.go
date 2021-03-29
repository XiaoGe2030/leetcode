package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{4, 5, 6, 1, 2, 3}))
}

func sortArray(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		min := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min] > nums[j] {
				min = j
			}
		}

		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}
