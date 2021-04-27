package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}

func rob(nums []int) int {
	/*
		还是动态规划
		dp[i] :表示索引0~i所在房子被偷时，的最大值
		动态方程：dp[i]= max{nums[i]+dp[i-2],dp[i-1]}
		初始状态: dp[0]=nums[0]
	*/

	n := len(nums)
	if n == 0 {
		return 0
	} else if n == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(nums[i]+dp[i-2], dp[i-1])
	}
	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
