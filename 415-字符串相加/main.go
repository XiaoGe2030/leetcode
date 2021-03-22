package main

import "fmt"

func main() {
	fmt.Println(addStrings("12234", "94219"))
}

func addStrings(num1 string, num2 string) string {

	i, j := len(num1)-1, len(num2)-1
	var isJ bool
	var res string
	for i >= 0 && j >= 0 {

		tmp := num1[i] - '0' + num2[j] - '0'
		if isJ {
			tmp++
		}

		if tmp >= 10 {
			isJ = true
			tmp -= 10
		} else {
			isJ = false
		}

		res += string(tmp + '0')
		i--
		j--
	}

	for ; i >= 0; i-- {
		tmp := num1[i] - '0'
		if isJ {
			tmp++
		}

		if tmp >= 10 {
			isJ = true
			tmp -= 10
		} else {
			isJ = false
		}
		res += string(tmp + '0')
	}

	for ; j >= 0; j-- {
		tmp := num2[j] - '0'
		if isJ {
			tmp++
		}

		if tmp >= 10 {
			isJ = true
			tmp -= 10
		} else {
			isJ = false
		}
		res += string(tmp + '0')
	}

	if isJ {
		res += string('1')
	}

	result := []rune(res)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}
