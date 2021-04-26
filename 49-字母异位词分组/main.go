package main

import "fmt"

func main() {
}

func groupAnagrams(strs []string) [][]string {
	/*
		map
		key：数组长度为26的
		value: 为string
	*/
	mp1 := make(map[[26]int][]string)
	fmt.Println(mp1)
	mp := map[[26]int][]string{}
	for _, str := range strs {
		cnt := [26]int{}
		for _, b := range str {
			cnt[b-'a']++
		}
		mp[cnt] = append(mp[cnt], str)
	}
	ans := make([][]string, 0, len(mp))
	for _, v := range mp {
		ans = append(ans, v)
	}
	return ans
}
