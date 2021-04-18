package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(1, 1))
	fmt.Println(uniquePaths(3, 7))
	fmt.Println(uniquePaths(7, 3))
	fmt.Println(uniquePaths(23, 12))
}

/*
	采用深度优先遍历吧
*/

func uniquePaths(m int, n int) int {
	var res int
	dfs(m, n, 0, 0, &res)
	return res
}

func dfs(m, n, x, y int, res *int) {
	if x == m-1 && y == n-1 {
		*res++
		return
	}

	//往右走一步
	if y < n-1 {
		dfs(m, n, x, y+1, res)
	}
	//下走一步
	if x < m-1 {
		dfs(m, n, x+1, y, res)
	}
}
