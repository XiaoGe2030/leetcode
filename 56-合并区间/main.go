package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func merge(intervals [][]int) [][]int {
	var res [][]int
	//把最小的左值放进res中
	//	bubb(intervals)

	sort.Sort(TwoArray(intervals))
	res = append(res, intervals[0])

	for i := 1; i < len(intervals); i++ {
		isAdd := false
		for j := 0; j < len(res); j++ {
			if checkIs(intervals[i][0], intervals[i][1], res[j][0], res[j][1]) {
				if intervals[i][0] <= res[j][0] {
					res[j][0] = intervals[i][0]
				}
				if intervals[i][1] >= res[j][1] {
					res[j][1] = intervals[i][1]
				}
				isAdd = true
			}
		}
		if !isAdd {
			res = append(res, intervals[i])
		}
	}
	return res
}

func checkIs(x1, y1, x2, y2 int) bool {
	if (y1 >= x2 && y1 <= y2) || (y2 >= x1 && y2 <= y1) {
		return true
	}

	return false
}

func bubb(res [][]int) {
	for i := len(res) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if res[j][0] > res[j+1][0] {
				res[j], res[j+1] = res[j+1], res[j]
			}
		}
	}
}

type TwoArray [][]int

func (ta TwoArray) Len() int           { return len(ta) }
func (ta TwoArray) Less(i, j int) bool { return ta[i][0] < ta[j][0] }
func (ta TwoArray) Swap(i, j int)      { ta[i], ta[j] = ta[j], ta[i] }
