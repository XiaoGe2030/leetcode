package main

import "fmt"

func main() {
	//fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	strMap := make(map[byte]int)
	var maxLen int

	var rightIndex int

	length := len(s)
	for i := 0; i < length; i++ {
		//滑动窗口左侧收缩,注意这里是删除i-1的索引
		if i != 0 {
			delete(strMap, s[i-1])
		}

		//右侧指定索引没有出现重复，索引就一直往右偏移
		for rightIndex < length && strMap[s[rightIndex]] == 0 {
			strMap[s[rightIndex]]++
			rightIndex++
		}

		//rightIndex索引位置不满足要求，所以从rightIndex-1开始计算
		if maxLen < rightIndex-i {
			maxLen = rightIndex - i
		}
	}
	return maxLen
}
