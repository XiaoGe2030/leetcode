package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	fmt.Println(reverse(23))
	fmt.Println(reverse(-23))
}

func reverse(x int) int {
	strX := strconv.Itoa(x)

	var r []rune

	isPos := 1

	if x < 0 {
		isPos = -1
		r = []rune(strX[1:])
	} else {
		r = []rune(strX)
	}

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	result, err := strconv.Atoi(string(r))

	if err != nil || result > math.MaxInt32 {
		return 0
	}
	return result * isPos
}
