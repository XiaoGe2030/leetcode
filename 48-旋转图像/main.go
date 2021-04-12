package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func rotate(matrix [][]int) {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}

	for i, row := range matrix {
		for j, v := range row {
			tmp[j][n-1-i] = v
		}
	}

	copy(matrix, tmp)
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[i][j], matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1] =
				matrix[n-j-1][i], matrix[n-i-1][n-j-1], matrix[j][n-i-1], matrix[i][j]
		}
	}
}

/*
	主要是找规律，
	顺时针旋转90°，
	[i,j]->[j,n-i-1]
	[j,n-i-1]->[n-i-1,n-j-1]
	[n-i-1,n-j-1]->[n-j-1,i]
	[n-j-1,i]->[i,j]
*/
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < (n+1)/2; j++ {
			matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i], matrix[i][j] =
				matrix[i][j], matrix[j][n-i-1], matrix[n-i-1][n-j-1], matrix[n-j-1][i]
		}
	}
}
