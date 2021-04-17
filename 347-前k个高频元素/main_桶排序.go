package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func topKFrequent(nums []int, k int) []int {
	var res []int
	resMap := make(map[int]int)
	for _, n := range nums {
		resMap[n]++
	}

	res2 := make([][]int, len(nums)+1)

	for k, v := range resMap {
		if res2[v] == nil {
			res2[v] = make([]int, 0)
		}
		res2[v] = append(res2[v], k)
	}

	for i := len(res2) - 1; i >= 0 && len(res) < k; i-- {
		if res2[i] == nil {
			continue
		}
		res = append(res, res2[i]...)
	}

	return res
}
