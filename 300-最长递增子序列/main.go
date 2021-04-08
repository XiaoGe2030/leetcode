package main

import "fmt"

func main() {
	fmt.Println(lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//1、确定状态，dp[i]表示以nums[i]结尾最长子序列的长度
	dp := make([]int, len(nums))
	dp[0] = 1 //2、初始条件
	maxans := 1
	//3、计算顺序
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			//状态转移方程
			if nums[i] > nums[j] {
				dp[i] = maxInt(dp[i], dp[j]+1)
			}
		}
		maxans = maxInt(maxans, dp[i])
	}
	return maxans
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
