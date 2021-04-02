package main

import (
	"fmt"
	"sort"
)

func main() {
	threeSumClosest([]int{-1, 2, 1, -4}, 2)
}

func threeSumClosest(nums []int, target int) [][]int {
	var res [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target_1 := target - nums[i]
		last := n - 1
		for second := i + 1; second < n; second++ {
			if second > i && nums[second] == nums[second-1] {
				continue
			}

			for second < last && nums[second]+nums[last] < target_1 {
				last--
			}

			if second == last {
				break
			}

			if nums[second]+nums[last] == target_1 {
				res = append(res, []int{nums[i], nums[second], nums[last]})
			}
		}
	}

	fmt.Println(res)
	return res
}
