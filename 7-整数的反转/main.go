package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(1534236469))
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	var res string
	if x < 0 {
		res = string(reverseStr([]rune(str[1:])))
	} else {
		res = string(reverseStr([]rune(str)))
	}

	result, _ := strconv.Atoi(res)
	if x < 0 {
		result = -result
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		result = 0
	}
	return result
}

func reverseStr(x []rune) []rune {
	for i, j := 0, len(x)-1; i < j; i, j = i+1, j-1 {
		x[i], x[j] = x[j], x[i]
	}

	return x
}
