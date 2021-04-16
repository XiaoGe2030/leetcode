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
	sByte := []byte(s)
	n := len(s)

	if n <= numRows || numRows <= 1 {
		return s
	}
	var resTmp [][]byte
	isCon := false
	for index := 0; index < n; index += numRows {
		if isCon {
			tmp := make([]byte, numRows)
			if index+numRows-2 > n {
				resver(tmp[1:numRows-1], sByte[index:])
			} else {
				resver(tmp[1:numRows-1], sByte[index:index+numRows-2])
			}
			resTmp = append(resTmp, tmp)
			index -= 2
		} else {
			if index+numRows > n {
				resTmp = append(resTmp, sByte[index:])
			} else {
				resTmp = append(resTmp, sByte[index:index+numRows])
			}
		}
		isCon = !isCon
	}

	var res []byte
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(resTmp); j++ {
			if i < len(resTmp[j]) {
				if resTmp[j][i] == 0 {
					continue
				} else {
					res = append(res, resTmp[j][i])
				}
			}
		}
	}
	return string(res)
}

func resver(t, s []byte) {
	for i := 0; i < len(s); i++ {
		t[len(t)-1-i] = s[i]
	}
}
