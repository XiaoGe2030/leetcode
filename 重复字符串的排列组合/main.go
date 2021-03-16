package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("qqe"))
}

func permutation(str string) []string {
	if len(str) == 0 {
		return nil
	}
	newStr := []string{}
	// 将字符串转换成字符串切片，方便排序
	for _, value := range str {
		newStr = append(newStr, string(value))
	}
	result := []string{}
	sort.Strings(newStr)
	dfsPermutation(newStr, 0, &result)
	return result
}

// 回溯函数实现
// i表示本次函数需要放置的元素位置
func dfsPermutation(str []string, i int, result *[]string) {
	n := len(str)
	if i == n-1 {
		var tmp string
		for _, value := range str {
			tmp += value
		}
		*result = append(*result, tmp)
		fmt.Println("result", *result)
		return
	}
	// nums[0:i]是已经决定的部分，nums[i:]是待决定部分，同时待选元素也都在nums[i:]
	for k := i; k < n; k++ {
		// 跳过重复数字
		if k != i && str[k] == str[i] {
			continue
		}
		str[k], str[i] = str[i], str[k]
		fmt.Println(str)
		dfsPermutation(str, i+1, result)
	}
	// 还原状态
	for k := n - 1; k > i; k-- {
		str[i], str[k] = str[k], str[i]
		fmt.Println(str)
	}
}
