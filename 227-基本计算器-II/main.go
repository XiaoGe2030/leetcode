package main

import "fmt"

func main() {
	fmt.Println(calculate("3+2*2"))
}

func calculate(s string) int {
	numStack := make([]int, 0)
	preSige := '+'
	num := 0
	for i, ch := range s {
		isDig := ch >= '0' && ch <= '9'
		if isDig {
			num *= 10
			num += int(ch - '0')
		}

		if !isDig && ch != ' ' || len(s)-1 == i {
			switch preSige {
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
			preSige = ch
		}
	}

	var res int
	for _, n := range numStack {
		res += n
	}

	return res

}
