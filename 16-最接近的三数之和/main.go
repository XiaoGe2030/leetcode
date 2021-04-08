package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 2))
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	min := math.MaxFloat64
	var res int

	for first := 0; first < n; first++ {

		second := first + 1
		third := n - 1
		for second < third {
			tmp := nums[first] + nums[second] + nums[third]
			absValues := math.Abs(float64(tmp - target))
			if tmp > target {
				third--
			} else if tmp < target {
				second++
			} else {
				return tmp
			}

			if absValues < min {
				res = tmp
				min = absValues
			}
		}

	}

	return res
}
