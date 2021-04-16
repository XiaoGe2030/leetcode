package main

import (
	"fmt"
)

func main() {
	//	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cat", "dog", "sand", "and", "catsan"}))
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true
	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			if dp[i] && sInSlice(s[i:j], wordDict) {
				dp[j] = true
			}
		}
	}

	return dp[n]
}

func sInSlice(s string, str []string) bool {
	for _, si := range str {
		if si == s {
			return true
		}
	}

	return false
}
