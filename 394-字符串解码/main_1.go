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
	var stackn []int
	var stacks []string
	for _, c := range s {
		if c == '[' {
			stackn = append(stackn, multi)
			stacks = append(stacks, res)
			multi, res = 0, ""
		} else if c == ']' {
			res = stacks[len(stacks)-1] + strings.Repeat(res, stackn[len(stackn)-1])
			stackn = stackn[:len(stackn)-1]
			stacks = stacks[:len(stacks)-1]
		} else if c >= '0' && c <= '9' {
			multi *= 10
			multi += int(c - '0')
		} else {
			res += string(c)
		}
	}
	return res
}
