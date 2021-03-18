package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbbd"))
}

func longestPalindrome(s string) string {
	s1 := s
	for i, j := 0, len(s)-1; i <= j; {
		if s[i] == s[j] {
			i++
			j--
		} else {
			s1 = s[i+1 : j]
		}
	}

	if len(s1) == 0 {
		s1 = s[:1]
	}

	return s1
}
