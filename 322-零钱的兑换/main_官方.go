package main

import (
	"fmt"
)

func main() {
	fmt.Println(coinChange([]int{2, 5, 7}, 27))
}

func coinChange(coins []int, amount int) int {
	max := amount + 1 //设置边界
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		//初始化设置边界情况
		dp[i] = max
	}
	//初始条件
	dp[0] = 0

	//算法时间复杂度，m*n
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			//计算顺序:从小到大
			if coins[j] <= i {
				dp[i] = minInt(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	fmt.Println(dp)
	return dp[amount]
}

func minInt(a, b int) int {
	if a > b {
		return b
	}

	return a
}
