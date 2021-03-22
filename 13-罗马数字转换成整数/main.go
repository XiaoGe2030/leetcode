package main

import "fmt"

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("IX"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}

func romanToInt(s string) int {
	var res int
	n := len(s)
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'I':
			if i+1 < n && s[i+1] == 'V' {
				res += 4
				i++
			} else if i+1 < n && s[i+1] == 'X' {
				res += 9
				i++
			} else {
				res++
			}
		case 'V':
			res += 5
		case 'X':
			if i+1 < n && s[i+1] == 'L' {
				res += 40
				i++
			} else if i+1 < n && s[i+1] == 'C' {
				res += 90
				i++
			} else {
				res += 10
			}
		case 'L':
			res += 50
		case 'C':
			if i+1 < n && s[i+1] == 'D' {
				res += 400
				i++
			} else if i+1 < n && s[i+1] == 'M' {
				res += 900
				i++
			} else {
				res += 100
			}
		case 'D':
			res += 500
		case 'M':
			res += 1000
		default:
			panic("no match case")
		}
	}
	return res
}
