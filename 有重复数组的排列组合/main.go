package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(permutation("qqe"))
}

func permutation(S string) []string {

	var tmpS []string
	for _, v := range S {
		tmpS = append(tmpS, string(v))
	}
	sort.Strings(tmpS)
	var res []string
	dfs(tmpS, 0, &res)
	return res
}

func dfs(str []string, i int, res *[]string) {
	//终止条件
	n := len(str)
	if i == n-1 {
		var tmp string
		for _, value := range str {
			tmp += value
		}
		*res = append(*res, tmp)
		return
	}
	//dfs处理
	for k := i; k < n; k++ {
		if k != i && str[i] == str[k] {
			continue
		}
		str[i], str[k] = str[k], str[i]
		dfs(str, i+1, res)
	}
	//状态回退
	for k := n - 1; k > i; k-- {
		str[i], str[k] = str[k], str[i]
	}
}
