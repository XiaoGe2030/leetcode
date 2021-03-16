package main

import (
	"fmt"
)

func main() {
	fmt.Println(permutation("qqe"))
}

var res []string

func permutation(S string) []string {
	res = nil
	maap := make(map[byte]int)
	//判断S中每个字母存在几次？
	for i, _ := range S {
		maap[S[i]] += 1
	}
	ll := make([]byte, 0)
	dfs(ll, maap, []byte(S))
	return res
}
func dfs(chance []byte, maap map[byte]int, s []byte) {
	//先来一个终止条件
	if len(chance) == len(s) {
		ll := make([]byte, len(chance))
		copy(ll, chance)
		res = append(res, string(ll))
		return
	}
	temp := make(map[byte]bool) //定义一个map 记录有哪些字母当前已经使用过
	for i, _ := range s {
		//根据map中是否存在元素，且temp中是否被使用，来标识
		if maap[s[i]] == 0 || temp[s[i]] == true {
			continue
		}
		chance = append(chance, s[i])
		temp[s[i]] = true
		maap[s[i]] -= 1
		dfs(chance, maap, s)
		chance = chance[:len(chance)-1]

		maap[s[i]] += 1
	}
}
