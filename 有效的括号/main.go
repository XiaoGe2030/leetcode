package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
}

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
func isValid(s string) bool {

	//手动实现一个栈，左括号进栈，右括号不进，判断栈顶元素是否匹配，匹配就弹出，最后确定栈的长度是否为0
	var str []rune
	for _, n := range s {
		if len(str) == 0 || n == '{' || n == '[' || n == '(' {
			str = append(str, n)
		}

		switch n {
		case ']':
			if str[len(str)-1] == '[' {
				str = str[:len(str)-1]
			} else {
				return false
			}
		case '}':
			if str[len(str)-1] == '{' {
				str = str[:len(str)-1]
			} else {
				return false
			}
		case ')':
			if str[len(str)-1] == '(' {
				str = str[:len(str)-1]
			} else {
				return false
			}
		}
	}

	if len(str) == 0 {
		return true
	}

	return false
}
