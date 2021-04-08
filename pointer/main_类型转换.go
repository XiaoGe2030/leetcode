package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	fmt.Printf("%#x  %#b \n", Float64bits(11.3), Float64bits(4)) // "0x3ff0000000000000"
	var intA int = 99
	uintA := Uint(intA)
	fmt.Printf("%#v %v  %v \n", intA, reflect.TypeOf(uintA), uintA)
}

//通过将float64类型指针转化为uint64类型指针，我们可以查看一个浮点数变量的位模式。
func Float64bits(f float64) uint64 {
	fmt.Println(reflect.TypeOf(unsafe.Pointer(&f)))            //unsafe.Pointer
	fmt.Println(reflect.TypeOf((*uint64)(unsafe.Pointer(&f)))) //*uint64
	return *(*uint64)(unsafe.Pointer(&f))
}
func Uint(i int) uint {
	return *(*uint)(unsafe.Pointer(&i))
}
