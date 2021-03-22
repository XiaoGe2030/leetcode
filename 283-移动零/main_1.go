package main

import "fmt"

func main() {
	fmt.Println([]int{0, 1, 0, 3, 12})
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}

func moveZeroes(nums []int) []int {
	//双指针法
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
	return nums
}
