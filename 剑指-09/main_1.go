package main

import "fmt"

func main() {
	fmt.Println("vim-go")
}

type CQueue struct {
	data   []int
	length int
}

func Constructor() CQueue {
	d := make([]int, 0)
	l := 0
	return CQueue{data: d, length: l}
}

func (this *CQueue) AppendTail(value int) {
	this.data = append(this.data, value)
	this.length++
}

func (this *CQueue) DeleteHead() int {
	if this.length == 0 {
		return -1
	}
	res := this.data[0]
	this.length--
	this.data = this.data[1:]
	return res
}
