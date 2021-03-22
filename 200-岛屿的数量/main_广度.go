package main

import "fmt"

func main() {
	fmt.Println(numIslands([][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}))

	fmt.Println(numIslands([][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}))
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}

	nr := len(grid)
	nc := len(grid[0])
	num_islands := 0

	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if grid[r][c] == '1' {
				num_islands++
				grid[r][c] = '0'
				tmp := make([]int, 0)
				tmp = append(tmp, r*nc+c)
				for len(tmp) > 0 {
					id := tmp[len(tmp)-1]
					tmp = tmp[:len(tmp)-1]
					row := id / nc
					col := id % nc
					if row-1 >= 0 && grid[row-1][col] == '1' {
						tmp = append(tmp, (row-1)*nc+col)
						grid[row-1][col] = '0'
					}
					if row+1 < nr && grid[row+1][col] == '1' {
						tmp = append(tmp, (row+1)*nc+col)
						grid[row+1][col] = '0'
					}
					if col-1 >= 0 && grid[row][col-1] == '1' {
						tmp = append(tmp, row*nc+col-1)
						grid[row][col-1] = '0'
					}
					if col+1 < nc && grid[row][col+1] == '1' {
						tmp = append(tmp, row*nc+col+1)
						grid[row][col+1] = '0'
					}
				}
			}
		}
	}

	return num_islands
}
