package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("  -42"))
}

func myAtoi(s string) int {
	var start int
	pos := true

	for start < len(s) {
		if s[start] == ' ' || s[start] != '0' {
			start++
		}
	}

	fmt.Println(start)
	if s[start] == '+' {
		start++
	} else if s[start] == '-' {
		start++
		pos = false
	}

	var res int
	fmt.Println(start)
	for i := start; i < len(s); i++ {
		switch s[i] {
		case '0':
			res *= 10
		case '1':
			res *= 10
			res += 1
		case '2':
			res *= 10
			res += 2
		case '3':
			res *= 10
			res += 3
		case '4':
			res *= 10
			res += 4
		case '5':
			res *= 10
			res += 5
		case '6':
			res *= 10
			res += 6
		case '7':
			res *= 10
			res += 7
		case '8':
			res *= 10
			res += 8
		case '9':
			res *= 10
			res += 9
		default:
			break
		}
	}

	if !pos {
		res = res * -1
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	} else if res < math.MinInt32 {
		return math.MinInt32
	}

	return res
}
