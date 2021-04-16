package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(removeDuplicateLetters("bcabc"))
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}

type myStr struct {
	Data []byte
}

func (m *myStr) Len() int           { return len(m.Data) }
func (m *myStr) Less(i, j int) bool { return m.Data[i] < m.Data[j] }
func (m *myStr) Swap(i, j int)      { m.Data[i], m.Data[j] = m.Data[j], m.Data[i] }
func removeDuplicateLetters(s string) string {

	m := &myStr{Data: []byte(s)}
	sort.Sort(m)
	var res []byte
	resMap := make(map[byte]byte)

	for _, c := range m.Data {
		if _, ok := resMap[c]; ok {
			continue
		} else {
			resMap[c] = c
			res = append(res, c)
		}
	}

	return string(res)
}
