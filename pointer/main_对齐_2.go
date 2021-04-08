package main

import (
	"fmt"
	"unsafe"
)

type T struct {
	t1 byte
	t2 int32
	t3 int64
	t4 string
	t5 bool
}

func writeByPointer() {
	uint6 := &Uint6{}
	lowPointer := (*[2]byte)(unsafe.Pointer(uint6))
	*lowPointer = [2]byte{1, 2}
	//unsafe.Offsetof会计算padding后的偏移距离
	//必须将unsafe.Pointer转化成 uintptr类型才能进行指针的运算，uintptr 与 unsafe.Pointer 之间可以相互转换。
	highPointer := (*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(uint6)) + unsafe.Offsetof(uint6.high)))
	fmt.Printf("addr %x\n addr %x\n size %v\n size %v\n size %v\n align %v\n offset %v \n",
		uintptr(unsafe.Pointer(uint6)),
		uintptr(unsafe.Pointer(uint6))+unsafe.Sizeof(uint6.low),
		unsafe.Sizeof([2]byte{1, 2}),
		unsafe.Sizeof(uint6.low),
		unsafe.Sizeof(uint6.high),
		unsafe.Alignof(uint6.low),
		unsafe.Offsetof(uint6.high))
	*highPointer = uint32(9)
	//借助于 unsafe.Pointer，我们实现了像 C 语言中的指针偏移操作。可以看出，这种不安全的操作使得我们可以在任何地方直接访问结构体中未公开的成员，只要能得到这个结构体变量的地址。
	fmt.Printf("%+v %v %v %v \n", uint6, &uint6, &uint6.low[0], &uint6.high)
}

type Uint6 struct {
	low  [2]byte
	high uint32
}

func main() {
	t := &T{1, 2, 3, "", true}
	fmt.Println("sizeof :")
	fmt.Println(unsafe.Sizeof(*t))
	fmt.Println(unsafe.Sizeof(t.t1))
	fmt.Println(unsafe.Sizeof(t.t2))
	fmt.Println(unsafe.Sizeof(t.t3))
	fmt.Println(unsafe.Sizeof(t.t4))
	fmt.Println(unsafe.Sizeof(t.t5))

	fmt.Println(":", unsafe.Alignof(*t))
	fmt.Println(":", unsafe.Alignof(t.t1))
	fmt.Println(":", unsafe.Alignof(t.t2))
	fmt.Println(":", unsafe.Alignof(t.t3))
	fmt.Println(":", unsafe.Alignof(t.t4))
	fmt.Println(":", unsafe.Alignof(t.t5))
	//这里以0x0作为基准内存地址。打印出来总共占用40个字节。t.t1 为 char，对齐值为 1，0x0 % 1 == 0，从0x0开始，占用一个字节；t.t2 为 int32，对齐值为 4，0x4 % 4 == 0，从 0x4 开始，占用 4 个字节；t.t3 为 int64，对齐值为 8，0x8 % 8 == 0，从 0x8 开始，占用 8 个字节；t.t4 为 string，对齐值为 8，0x16 % 8 == 0，从 0x16 开始， 占用 16 个字节（string 内部实现是一个结构体，包含一个字节类型指针和一个整型的长度值）；t.t5 为 bool，对齐值为 1，0x32 % 8 == 0，从 0x32 开始，占用 1 个字节。从上面分析，可以知道 t 的对齐值为 8，最后 bool 之后会补齐到 8 的倍数，故总共是 40 个字节。

	fmt.Println("Offsetof : ")
	fmt.Println(unsafe.Offsetof(t.t1))
	fmt.Println(unsafe.Offsetof(t.t2))
	fmt.Println(unsafe.Offsetof(t.t3))
	fmt.Println(unsafe.Offsetof(t.t4))
	fmt.Println(unsafe.Offsetof(t.t5))

	writeByPointer()
	//CPU看待内存是以block为单位的，就像是linux下文件大小的单位IO block为4096一样，
	//是一种牺牲空间换取时间的做法, 我们一定要注意不要浪费空间，
	//struct类型定义的时候一定要将占用内从空间小的类型放在前面, 充足利用padding， 才能提升内存、cpu效率
}
