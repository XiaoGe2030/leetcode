package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("vim-go")
}

func topKFrequent(nums []int, k int) []int {
	nMap := make(map[int]int)
	var resK []int
	for _, v := range nums {
		nMap[v]++
	}

	for k := range nMap {
		resK = append(resK, k)
	}

	m := &myMap{Data: nMap, Res: resK}
	sort.Sort(m)

	return m.Res[0:k]
}

type myMap struct {
	Data map[int]int //map中key是数组中元素，value是元素个数
	Res  []int       //切片中是元素中不重复元素的个数
}

func (m *myMap) Len() int           { return len(m.Data) }
func (m *myMap) Less(i, j int) bool { return m.Data[m.Res[i]] > m.Data[m.Res[j]] }
func (m *myMap) Swap(i, j int)      { m.Res[i], m.Res[j] = m.Res[j], m.Res[i] }
