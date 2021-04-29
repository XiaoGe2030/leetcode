package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	var s int
	strconv.Itoa()
}

func multiply(num1 string, num2 string) string {
	n1, n2 := len(num1), len(num2)
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	res := make([]int, n1+n2)

	for i := n2 - 1; i >= 0; i-- {
		ni := int(num2[i] - '0')
		for j := n1 - 1; j >= 0; j-- {
			nj := int(num1[j] - '0')
			tmp := res[i+j+1] + ni*nj
			res[i+j+1] = tmp % 10
			res[i+j] += tmp / 10 //注意这里跟i+j+1处理的不一样，主要是多位数相乘时，下一次循环，上一次外层循环的值需要保留。
			/*
				123
				x    54
				---------
				492
				15
			*/
		}
	}

	var resStr string

	for i := 0; i < len(res); i++ {
		if i == 0 && res[i] == 0 {
			continue
		}

		resStr += strconv.Itoa(res[i])
	}

	return resStr
}
