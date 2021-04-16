package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

var res [][]string
var temp []string
var  n int

func isPalind(s []string) bool {
	for i,j := 0,len(s); i < j ; i,j=i+1,j-1 {
		if s[i] != s[j] {
		return false
		}
	}
	return true
}

func backtravel(string s, int index) {
	if index == n {
		res = append(res,temp)
		return
	}
	for i := index; i < n; i++ {
		if(isPalind(s[index:i-index+1])){
			temp= append(temp,s[index:i-index+1])
			backtravel(s, i+1)
			temp = temp[:len(temp)-1]
		}
	}
}

func partition( s string) [][]string {
	n = len(s) 
	if(n == 0)
	return res
	backtravel(s, 0)
	return res
}
