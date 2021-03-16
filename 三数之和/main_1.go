package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	fmt.Println(threeSum1([]int{-1, 0, 1, 2, -1, -4}))
}

//在数组中找到三个数满足a+b+c = 0
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	var res [][]int
	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		target := -nums[a]
		c := n - 1
		for b := a + 1; b < n; b++ {
			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			for b < c && nums[b]+nums[c] > target {
				c--
			}
			if b == c {
				break
			}
			if nums[b]+nums[c] == target {
				res = append(res, []int{nums[a], nums[b], nums[c]})
			}
		}
	}
	return res
}

func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	for a := 0; a < n; a++ {
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}
		b := a + 1
		c := n - 1
		target := -nums[a]

		for b < c {
			sum := nums[c] + nums[b]
			if sum == target {
				res = append(res, []int{nums[a], nums[b], nums[c]})
				for b < c && nums[b] == nums[b+1] {
					b++
				}
				b++

				for b < c && nums[c] == nums[c-1] {
					c--
				}
				c--

			} else if sum < target {
				b++
			} else {
				c--
			}
		}
	}
	return res
}
