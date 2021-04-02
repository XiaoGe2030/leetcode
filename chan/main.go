package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 2)
	fmt.Println("len chan :", len(ch))
	fmt.Println(^uintptr(0))
	go func(ch chan<- int) {
		ch <- 1
		fmt.Println("len chan :", len(ch))
		ch <- 2
		fmt.Println("len chan :", len(ch))
		ch <- 3
		fmt.Println("len chan :", len(ch))

		close(ch)
		fmt.Println("ch closed")
	}(ch)

	time.Sleep(100 * time.Second)
	for {
		select {
		case i := <-ch:
			fmt.Println("receive :", i)
			time.Sleep(time.Second)
		case <-time.After(time.Second):
			fmt.Println("time out")
			//os.Exit(1)
			return
		}
	}

	fmt.Println("Game Over")
}
