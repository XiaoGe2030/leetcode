package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

func partition(s string) [][]string {

}

func isCheck(str []string) bool {
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		if str[i] != str[j] {
			return false
		}
	}

	return true
}
