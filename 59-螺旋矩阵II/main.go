package main

import "fmt"

func main() {
	fmt.Println(generateMatrix(3))
	fmt.Println(generateMatrix(4))
	fmt.Println(generateMatrix(5))
}

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}

	m := n
	var num int
	for k := 0; n > 0; n-- {
		i, j := k, k
		for j < n-1 {
			num++
			res[i][j] = num
			j++
		}

		for i < n-1 {
			num++
			res[i][j] = num
			i++
		}

		for j > k {
			num++
			res[i][j] = num
			j--
		}

		for i > k {
			num++
			res[i][j] = num
			i--
		}
		k++
	}

	if m&1 == 1 {
		res[(m)/2][(m)/2] = m * m
	}
	return res

}
