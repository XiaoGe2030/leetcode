package main

import "fmt"

func main() {
	fmt.Println(addStrings("12234", "94219"))
}

func addStrings(num1 string, num2 string) string {

	i, j := len(num1)-1, len(num2)-1
	var res string
	var carry int
	for i >= 0 || j >= 0 {
		var n1, n2 int
		if i >= 0 {
			n1 = int(num1[i] - '0')
		}

		if j >= 0 {
			n2 = int(num2[j] - '0')
		}

		tmp := n1 + n2 + carry
		carry = tmp / 10
		res += string(tmp%10 + '0')
		i--
		j--
	}

	if carry == 1 {
		res += string('1')
	}
	result := []rune(res)
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)

}
