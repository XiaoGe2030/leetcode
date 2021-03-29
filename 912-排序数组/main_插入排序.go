package main

import "fmt"

func main() {
	//fmt.Println(sortArray([]int{1, 2, 4, 5, 6, 3}))
	fmt.Println(sortArray([]int{4, 5, 6, 1, 2, 3}))
}

func sortArray(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		tmp := nums[i+1]
		j := i
		for j >= 0 && tmp < nums[j] {
			nums[j+1] = nums[j]
			j--
		}

		nums[j+1] = tmp
	}
	return nums
}
