package main

import "fmt"

func main() {
	var str = "hello 你好"
	fmt.Println(len(str))

	fmt.Println(len([]rune(str)))
}
