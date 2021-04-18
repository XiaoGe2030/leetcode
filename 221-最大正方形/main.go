package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func maximalSquare(matrix [][]byte) int {
	/*
		动态规划
		dp[i,j]表示[i,j]为正方形的左上角节点时的最大正方形面积
		初始化:
		dp[i,n-1]=matrix[i,n-1]
		dp[m-1,j]=matrix[m-1,j]
		状态转移方程：
		如果matrix[i,j]=0,则dp[i,j]=0
		if dp[i+1,j+1] = dp[i+1,j]=dp[i,j+1]
		则dp[i,j]= dp[i+1,j]+1
		else
		dp[i,j]=min{dp[i+1,j+1],dp[i+1,j],dp[i,j+1]}+1
	*/
	max := 0
	m := len(matrix)
	n := len(matrix[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if matrix[i][n-1] == '1' {
			res[i][n-1] = 1
			max = 1
		}

	}

	for j := 0; j < n; j++ {
		if matrix[m-1][j] == '1' {
			res[m-1][j] = 1
			max = 1
		}
	}

	for i := m - 2; i >= 0; i-- {
		for j := n - 2; j >= 0; j-- {
			if matrix[i][j] == '1' {
				minV := min(min(res[i+1][j], res[i][j+1]), res[i+1][j+1])
				res[i][j] = minV + 1
				if max < minV+1 {
					max = minV + 1
				}
			}
		}
	}

	return max * max

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
