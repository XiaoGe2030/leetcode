package main

import "fmt"

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
	fmt.Println(letterCombinations("2"))
}

var numMap = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

var res []string

func letterCombinations(digits string) []string {

	res = res[:0]
	var str string
	huisu(digits, 0, str)
	return res
}

func huisu(digits string, index int, str string) {
	if len(digits) == len(str) {
		res = append(res, str)
	} else {
		letters := numMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			huisu(digits, index+1, str+string(letters[i]))
		}
	}
}
