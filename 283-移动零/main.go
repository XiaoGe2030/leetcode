package main

import "fmt"

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}

func moveZeroes(nums []int) []int {

	tmp := make([]int, len(nums))
	for i, j := 0, 0; i < len(nums); i++ {
		if nums[i] != 0 {
			tmp[j] = nums[i]
			j++
		}
	}

	return tmp
}
