package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxInt32)
	fmt.Println(math.MinInt32)
	fmt.Println(myAtoi("9223372036854775808"))
	fmt.Println(myAtoi("-91283472332"))
	fmt.Println(myAtoi("-2147483647"))
}

func myAtoi(s string) int {
	var start int
	pos := true

	n := len(s)
	for start < n && s[start] == ' ' {
		start++
	}

	var res int
	for i := start; i < n; i++ {
		if i == start && s[i] == '+' {
			pos = true
		} else if i == start && s[i] == '-' {
			pos = false
		} else {
			if s[i] < '0' || s[i] > '9' {
				break
			}
			res *= 10
			res += int(s[i] - '0')
			if res > math.MaxInt32 {
				if pos {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			}
		}
	}

	if !pos {
		res = res * -1
	}

	return res
}
