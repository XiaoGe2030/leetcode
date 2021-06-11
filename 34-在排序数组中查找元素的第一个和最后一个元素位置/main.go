package main

import "fmt"

func main() {
	fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
}

func searchRange(nums []int, target int) []int {

	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2
		if nums[mid] == target {
			befor := mid
			after := mid

			for befor >= 0 && nums[befor] == target {
				befor--
			}

			for after <= len(nums)-1 && nums[after] == target {
				after++
			}
			return []int{befor + 1, after - 1}
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return []int{-1, -1}
}
