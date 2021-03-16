package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("babad"))
}

//动态规划方法，p[i][j]= (s[i]==s[j]) and (p[i+1][j-1])
func longestPalindrome(s string) string {
	n := len(s)
	resSlice := make([][]int, n)
	for i, _ := range s {
		resSlice[i] = make([]int, n)
	}

	var resStr string
	for k := 0; k < n; k++ {
		for i := 0; i < n-k; i++ {
			j := i + k
			if k == 0 {
				resSlice[i][j] = 1
			} else if k == 1 {
				if s[i] == s[j] {
					resSlice[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					resSlice[i][j] = resSlice[i+1][j-1]
				}
			}

			if resSlice[i][j] > 0 && k+1 > len(resStr) {
				resStr = s[i : j+1]
			}
		}
	}

	return resStr
}
