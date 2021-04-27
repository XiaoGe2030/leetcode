package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(removeKdigits("1432219", 3))
}

func removeKdigits(num string, k int) string {
	stack := make([]byte, 0)
	remain := len(num) - k
	for i, _ := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k -= 1
		}
		stack = append(stack, num[i])
	}

	stack = bytes.TrimLeft(stack[:remain], "0")
	if len(stack) == 0 {
		return "0"
	}

	return string(stack)
}
