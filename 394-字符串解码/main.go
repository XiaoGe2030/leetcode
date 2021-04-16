package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(decodeString("3[a]2[bc]"))
	fmt.Println(decodeString("3[a2[c]]"))
	fmt.Println(decodeString("2[abc]3[cd]ef"))
	fmt.Println(decodeString("abc3[cd]xyz"))
}

func decodeString(s string) string {
	multi, res := 0, ""

	var stackMulti []int
	var stackStr []string
	for _, c := range s {
		if c == '[' {
			stackMulti = append(stackMulti, multi)
			stackStr = append(stackStr, res)
			multi, res = 0, ""
		} else if c == ']' {
			cur := stackMulti[len(stackMulti)-1]
			stackMulti = stackMulti[:len(stackMulti)-1]
			tmp := strings.Repeat(res, cur)
			res = stackStr[len(stackStr)-1] + tmp
			stackStr = stackStr[:len(stackStr)-1]
		} else if c >= '0' && c <= '9' {
			multi *= 10
			multi += int(c - '0')
		} else {
			res += string(c)
		}
	}
	return res
}
