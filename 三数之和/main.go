package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		target := -1 * nums[first]
		third := n-1 
	}
}

func w
