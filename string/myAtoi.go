package main

import (
	"math"
	"strconv"
	"strings"
)

func main() {
	math.MaxInt32
	math.MinInt32

	strconv
	strings.TrimLeft()
}

func myAtoi(s string) int {

	var i int
	isPos := true
	for _, ch := range s {
		switch ch {
		case "":
			i++
		case "-":
			isPos = false
			i++
			break
		case "+":
			i++
			break
		default:
			break
		}
	}

	if i == len(s) {
		return 0
	}

	for j := i; j < len(s); j++ {
		if s[j] < '0' && s[j] > '9' {
			break
		}
	}

	if j == i {
		return 0
	}

	rStr := s[i:j]
	result, err := strconv.Atoi(rStr)
	if err != nil {

	}

	if !isPos {
		result = -result
	}
	return result

}
