package main

import (
	"fmt"
	"unsafe"
)

type W struct {
	b byte
	i int32
	j int64
}

func main() {
	var w W = W{}
	//在struct中，它的对齐值是它的成员中的最大对齐值。
	fmt.Printf("%v, %v, %v, %v, %v, %v, %v, %v\n",
		unsafe.Alignof(w),
		unsafe.Alignof(w.b),
		unsafe.Alignof(w.i),
		unsafe.Alignof(w.j),
		unsafe.Sizeof(w),
		unsafe.Sizeof(w.b),
		unsafe.Sizeof(w.i),
		unsafe.Sizeof(w.j))

	fmt.Println("byte :", unsafe.Alignof(byte(0)))
	fmt.Println("int8 :", unsafe.Alignof(int8(0)))
	fmt.Println("unit8 :", unsafe.Alignof(uint8(0)))
	fmt.Println("int16 :", unsafe.Alignof(int16(0)))
	fmt.Println("uint16 :", unsafe.Alignof(uint16(0)))
	fmt.Println("int32 :", unsafe.Alignof(int32(0)))
	fmt.Println("uint32 :", unsafe.Alignof(uint32(0)))
	fmt.Println("int64 :", unsafe.Alignof(int64(0)))
	fmt.Println("uint64 :", unsafe.Alignof(uint64(0)))
	fmt.Println("uintptr :", unsafe.Alignof(uintptr(0)))
	fmt.Println("float32 :", unsafe.Alignof(float32(0)))
	fmt.Println("float64 :", unsafe.Alignof(float64(0)))
	//fmt.Printl" :",n(unsafe.Alignof(complex(0, 0)))
	fmt.Println("complex64 :", unsafe.Alignof(complex64(0)))
	fmt.Println("complex128 :", unsafe.Alignof(complex128(0)))
	fmt.Println("kong :", unsafe.Alignof(""))
	fmt.Println("new :", unsafe.Alignof(new(int)))
	fmt.Println("struct :", unsafe.Alignof(struct {
		f  float32
		ff float64
	}{}))
	fmt.Println(unsafe.Alignof(make(chan bool, 10)))
	fmt.Println(unsafe.Alignof(make([]int, 10)))
	fmt.Println(unsafe.Alignof(make(map[string]string, 10)))
}
