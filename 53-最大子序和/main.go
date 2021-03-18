package main

import "fmt"

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}

func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		//这里求出索引i结尾的最大数
		if nums[i]+nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		//索引i为结尾的最大数，跟当前最大值得对比
		if nums[i] > max {
			max = nums[i]
		}
	}

	return max
}
