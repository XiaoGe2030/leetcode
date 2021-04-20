package main

import (
	"fmt"
)

func main() {
	//	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate("3/2"))
	//	fmt.Println(calculate("3+ 5 / 2"))
	//	fmt.Println(calculate("1-1+1"))
}

func calculate(s string) int {
	numStack := make([]int, 0)
	num := 0
	preSign := '+'
	for i, ch := range s {
		isDig := ch >= '0' && ch <= '9'
		if isDig {
			num *= 10
			num += int(ch - '0')
		}

		if !isDig && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				numStack = append(numStack, num)
			case '-':
				numStack = append(numStack, -num)
			case '*':
				numStack[len(numStack)-1] *= num
			case '/':
				numStack[len(numStack)-1] /= num
			}
			num = 0
			preSign = ch
		}
	}

	var res int
	for _, n := range numStack {
		res += n
	}

	return res
}
