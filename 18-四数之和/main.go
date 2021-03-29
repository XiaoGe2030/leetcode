package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		res = append(res, threeSum(nums[i+1:], target-nums[i], nums[i])...)
	}
	return res
}

func threeSum(nums []int, target, index int) [][]int {

	sort.Ints(nums)
	var res [][]int
	n := len(nums)
	for i := 0; i < n; i++ {
		//排除第一个参数重复的情况
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target_2 := target - nums[i]
		k := n - 1
		for j := i + 1; j < n; j++ {
			//排除第二个参数重复的情况
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for j < k && nums[j]+nums[k] > target_2 {
				k--
			}

			if j == k {
				break
			}

			if nums[j]+nums[k] == target_2 {
				res = append(res, []int{index, nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}
