package main

import "fmt"

func main() {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

//利用单调栈，
func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)

	//单调栈，存储T数组的下标
	var tmp []int
	for i, t := range T {
		if len(tmp) == 0 {
			tmp = append(tmp, i)
		} else {
			for len(tmp) > 0 {
				index := tmp[len(tmp)-1]
				if T[index] >= t {
					break
				}
				res[index] = i - index
				tmp = tmp[:len(tmp)-1]
			}

			tmp = append(tmp, i)
		}
	}

	return res
}
