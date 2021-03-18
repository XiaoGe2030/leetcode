package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

var result [][]int
var numBool []bool

func permute(nums []int) [][]int {
	tmp := make([]int, 0)
	numBool = make([]bool, len(nums))
	getNum(nums, 0, tmp)
	return result
}

func getNum(nums []int, index int, tmp []int) {

	if index == len(nums) {
		numTmp := make([]int, 0)
		numTmp = append(numTmp, tmp...)
		result = append(result, numTmp)
		return
	}

	for k, v := range nums {
		if !numBool[k] {
			tmp = append(tmp, v)
			numBool[k] = true
			getNum(nums, index+1, tmp)
			tmp = tmp[:len(tmp)-1]
			numBool[k] = false
		}
	}
	return
}
