package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
	sortInt([]int{3, 5, 1})
}

func sortInt(nums []int) []int {
	fmt.Println(nums)
	sort.Sort(sort.IntSlice(nums))
	fmt.Println(nums)
	return nums
}
