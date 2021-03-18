package main

import (
	"fmt"
	"sort"
)

//nil slice 也可以调用slice的内置函数
var resultSlice [][]int

func main() {
	getNumSlices([]int{2, 3, 1, 6, 7, 8, 3, 4, 5}, 7)

	fmt.Println(resultSlice)
}

func getNumSlices(nums []int, m int) [][]int {

	fmt.Println(nums)
	sort.Ints(nums)
	fmt.Println(nums)

	if len(nums) == 0 {
		return nil
	}

	resultSlice = append(resultSlice, (getNums(nums, m)))

	return nil
}

func getNums(nums []int, m int) []int {
	var tmpSlice []int
	var i int
	for ; i < len(nums) && m >= 0; i++ {
		if nums[i] <= m {
			tmpSlice = append(tmpSlice, nums[i])
			m -= nums[i]
		}
	}
	if m == 0 {
		return tmpSlice
	} else {
		//切片中的值，先弹出，在把和减去
		if len(tmpSlice) != 0 {
			tmpSlice = tmpSlice[:len(tmpSlice)-1]
			m += nums[i-1]

		}
		tmpSlice = append(tmpSlice, getNums(nums[i:], m)...)
		return tmpSlice
	}
}
