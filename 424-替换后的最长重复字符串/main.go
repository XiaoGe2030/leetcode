package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
}

func characterReplacement(s string, k int) int {
	cnt := [26]int{}
	maxCnt, left := 0, 0
	for right, ch := range s {
		cnt[ch-'A']++
		//maxCnt记录的是数组中出现次数最多的元素个数
		maxCnt = max(maxCnt, cnt[ch-'A']) //maxCnt不必每次更新，这里有点记录最长重复字符串的意思
		//判断出不重复的值得个数如果大于k，则左边界的元素需要被移除
		if right-left+1-maxCnt > k { //这里其实maxCnt+k，维护了滑动窗口的最小值，所以k固定，所以maxCnt只会需要记录增长值就行
			cnt[s[left]-'A']--
			left++
		}
	}
	//fmt.Println(string(s[left:]))//s[left:]的结果不一定是重复字母最长子串，但是长度是等于结果预期的
	//return len(s) - left
	return min(maxCnt+k, len(s)) //这两个返回值都对
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
