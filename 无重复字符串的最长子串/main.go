package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	//fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("dvdl"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	n := len(s)

	rk, res := -1, 0

	for i := 0; i < n; i++ {
		if i != 0 {
			delete(m, s[i-1])
		}

		for rk+1 < n && m[s[rk+1]] == 0 {
			m[s[rk+1]]++
			rk++
		}
		res = max(res, rk-i+1)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

/*暴力求解法不满足，时间复杂度要求
func lengthOfLongestSubstring(s string) int {

	var maxLen int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			tmp := j - i + 1
			if tmp > maxLen && !isRepeat(s[i:j+1]) {
				maxLen = tmp
			}
		}
	}

	return maxLen
}

func isRepeat(str string) bool {
	resMap := make(map[rune]int)

	for _, s := range str {
		if _, ok := resMap[s]; ok {
			return true
		}
		resMap[s] = 1
	}

	return false
}
*/
