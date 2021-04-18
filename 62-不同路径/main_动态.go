package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(1, 1))
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(23, 12))
}

func uniquePaths(m int, n int) int {
	//深度优先遍历超时
	//动态规划
	//动态方程
	/*
		动态方程：dp[i,j]=dp[i+1,j]+dp[i,j+1]
		初始条件：dp[i,n-1]=1 (i<m-1)
		dp[m-1,j]=1 (j<n-1)
		计算方向：大 -> 小
	*/
	res := make([][]int, 0)
	for i := 0; i < m; i++ {
		res = append(res, make([]int, n))
	}

	for i := 0; i <= m-1; i++ {
		res[i][n-1] = 1
	}

	for j := 0; j < n-1; j++ {
		res[m-1][j] = 1
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			res[i][j] = res[i+1][j] + res[i][j+1]
		}
	}

	return res[0][0]
}
