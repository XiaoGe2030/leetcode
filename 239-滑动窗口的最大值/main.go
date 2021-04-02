package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}

//暴力解法
func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	for i := 0; i <= len(nums)-k; i++ {
		tmp := make([]int, 0)
		tmp = append(tmp, nums[i:i+k]...)
		sort.Ints(tmp)

		res = append(res, tmp[k-1])
	}

	return res
}
