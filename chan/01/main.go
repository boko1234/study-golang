package main

import (
	"fmt"
	"sync"
)

var a []int
var b chan int
var wg sync.WaitGroup

func fa() {
	fmt.Println(b)
	b = make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("10发送到b通道中。。。", x)
	}()
	//<- b				//接收
	b <- 10 //卡住了
	fmt.Println("10发送到b通道中了")
	b = make(chan int, 16)
	fmt.Println(b)
	wg.Wait()
}

func fb() {
	fmt.Println(b)
	b = make(chan int, 10)
	b <- 10
	fmt.Println("发送到通道中了")
	b <- 20
	fmt.Println("发送20中")
	x := <-b
	fmt.Println("从通道中取到了",x)
}

func main() {
	fb()
}
