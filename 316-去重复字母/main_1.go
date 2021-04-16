package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}

/*
	1、都为小写字母，可以用长度为26的数组表示,数组表示剩余的元素
	2、用一个栈存储已经存在的元素
	3、在用一个数组，表示栈中的元素，
*/
func removeDuplicateLetters(s string) string {
	var left [26]int
	for _, c := range s {
		left[c-'a']++
	}

	var stack []byte
	var inStack [26]bool

	for i := range s {
		c := s[i]
		if !inStack[c-'a'] {
			for n := len(stack); n > 0 && c < stack[n-1]; n = len(stack) {
				last := stack[n-1] - 'a'
				if left[last] == 0 {
					break
				}
				stack = stack[:n-1]
				inStack[last] = false
			}
			stack = append(stack, c)
			inStack[c-'a'] = true
		}
		left[c-'a']--
	}

	return string(stack)
}
