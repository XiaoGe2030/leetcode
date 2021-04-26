package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func removeDuplicates(S string) string {
	/*
		用一个栈来处理
	*/

	var stack []byte
	for i := 0; i < len(S); i++ {
		n := len(stack)
		if n > 0 && stack[n-1] == S[i] {
			stack = stack[:n-1]
		} else {
			stack = append(stack, S[i])
		}
	}

	return string(stack)
}
