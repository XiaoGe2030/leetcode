package main

import (
	"fmt"
)

func main() {
	fmt.Println("PAYPALISHIRING")
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println("PAHNAPLSIIGYIR")
}

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rows := make([]string, numRows)

	i, flag := 0, -1
	for _, c := range s {
		rows[i] = append(rows[i], c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	res := ""
	for _, r := range rows {
		res += r
	}

	return res
}
