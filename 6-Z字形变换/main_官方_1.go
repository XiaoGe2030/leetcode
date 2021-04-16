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
	if numRows == 1 {
		return s
	}

	rows := make([]string, min(numRows, len(s)))

	curRow := 0
	goingDown := false

	for _, c := range s {
		rows[curRow] += string(c)
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	res := ""
	for _, row := range rows {
		res += row
	}

	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
