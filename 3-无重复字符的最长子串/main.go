package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("ab"))
}

func lengthOfLongestSubstring(s string) int {

	strMap := make(map[byte]int)
	var maxLen, rightIndex int
	n := len(s)
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(strMap, s[i-1])
		}

		for rightIndex < n && strMap[s[rightIndex]] == 0 {
			strMap[s[rightIndex]]++
			rightIndex++
		}

		if maxLen < rightIndex-i {
			maxLen = rightIndex - i
		}
	}
	return maxLen
}
