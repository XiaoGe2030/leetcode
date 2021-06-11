package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(numRabbits([]int{1, 1, 2}))
}

func numRabbits(answers []int) int {
	resMap := make(map[int]int)
	res := 0
	for _, i := range answers {
		resMap[i]++
	}

	for i, v := range resMap {
		//需要向上取整
		res += math.Ceil(v/(i+1)) * (i + 1)
	}
	return res
}
