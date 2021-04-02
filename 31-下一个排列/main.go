package main

import (
	"fmt"
	"sort"
)

func main() {
	nextPermutation([]int{1, 2, 3, 8, 5, 7, 6, 4})
	nextPermutation([]int{3, 2, 1})
}

func nextPermutation(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}

	last := n - 1
	for ; last > 0; last-- {
		if nums[last] > nums[last-1] {
			break
		}
	}

	if last > 0 {
		for i := n - 1; i >= last; i-- {
			if nums[last-1] < nums[i] {
				nums[last-1], nums[i] = nums[i], nums[last-1]
				break
			}
		}
	}

	sort.Ints(nums[last:])

	fmt.Println(nums)
}
