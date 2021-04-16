package main

import "fmt"

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
	fmt.Println(wordBreak("applepenapple", []string{"apple", "pen"}))
	fmt.Println(wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
}

func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := i + 1; j < n+1; j++ {
			if dp[i] && inStr(s[i:j], wordDict) {
				dp[j] = true
			}
		}
	}

	return dp[n]
}

func inStr(s string, str []string) bool {
	for _, si := range str {
		if s == si {
			return true
		}
	}

	return false
}
