package main

import (
	"fmt"
	"math/big"
)

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
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
