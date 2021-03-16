package main

import (
	"fmt"
)

func main() {
	fmt.Println(permutation("qqe"))
}

func permutation(S string) []string {
	if len(S) == 0 {
		return nil
	}

	var newStr []string
	var result []string
	for _, value := range S {
		newStr = append(newStr, string(value))
	}

	//sort.Strings(newStr)
	dfs(newStr, 0, &result)
	return result
}

func dfs(str []string, i int, res *[]string) {

	n := len(str)
	//回溯的终止条件
	if n-1 == i {
		var tmp string
		for _, v := range str {
			tmp += v
		}
		*res = append(*res, tmp)
		return
	}
	//dfs处理
	for k := i; k < n; k++ {
		if k != i && str[k] == str[i] {
			continue
		}

		str[i], str[k] = str[k], str[i]
		dfs(str, i+1, res)
	}

	//状态太回滚
	for j := n - 1; j >= i; j-- {
		str[i], str[j] = str[j], str[i]
	}
}
