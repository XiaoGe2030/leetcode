package main

import (
	"fmt"
)

func main() {
	fmt.Println(findRepeatNumber([]int{2, 3, 1, 0, 2, 5, 3}))
	fmt.Println(findRepeatNumber([]int{3, 1, 2, 3}))
	fmt.Println(findRepeatNumber([]int{0, 1, 2, 3, 4, 11, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}))
}

func findRepeatNumber(nums []int) int {
	numMap := make(map[int]int)

	for _, n := range nums {
		if _, ok := numMap[n]; ok {
			numMap[n] = n
		} else {
			return n
		}
	}
	return -1
}
