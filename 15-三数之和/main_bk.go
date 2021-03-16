package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
	fmt.Println(threeSum1([]int{-2, 0, 1, 1, 2}))
}

func threeSum(nums []int) [][]int {

	var result [][]int
	sort.Ints(nums)
	n := len(nums)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		target := -nums[first]
		third := n - 1
		for second := first + 1; second < n; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			for second < third && nums[third]+nums[second] > target {
				third--
			}

			//if second == third {
			//	break
			//}

			if nums[third]+nums[second] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return result
}

func threeSum1(nums []int) [][]int {

	var result [][]int
	sort.Ints(nums)
	n := len(nums)

	for first := 0; first < n; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		second := first + 1
		third := n - 1
		target := -nums[first]

		for second < third {
			sum := nums[second] + nums[third]
			if sum == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
				for second < third && nums[second] == nums[second+1] {
					second++
				}
				second++

				for second < third && nums[third] == nums[third-1] {
					third--
				}

				third--
			} else if sum < target {
				second++
			} else {
				third--
			}
		}
	}
	return result
}
