package main

import "fmt"

func main() {
<<<<<<< HEAD
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
=======
	fmt.Println("vim-go")
}

func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
>>>>>>> 7f8765925abe64153e71a660dd25ace3debcd500
		}
	}

	return res
}
