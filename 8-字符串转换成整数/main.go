package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("  -42"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("+3-1"))
	fmt.Println(myAtoi("00000-42a1234"))
	fmt.Println(myAtoi("  -0012a42"))
}

func myAtoi(s string) int {
	var start int
	pos := true

	n := len(s)
	for start < n {
		if s[start] != ' ' {
			break
		}
		start++
	}

	var res int
	for i := start; i < n; i++ {
		if i == start {
			if s[i] == '+' {
				pos = true
			} else if s[i] == '-' {
				pos = false
			}
		} else {
			if s[i] < '0' || s[i] > '9' {
				return res
			}
			res *= 10
			res += int(s[i] - '0')
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
