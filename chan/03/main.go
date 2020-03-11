package main

import (
	"fmt"
	"sync"
)

/*//TODO://
//channel练习
//启动一个goroutine 生成100个数发送到
*/

var wg sync.WaitGroup

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

func f2(ch1, ch2 chan int) {
	defer wg.Done()
	//for x := range ch1 {
	//	ch2 <- x * x
	//}
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}

}

func main() {
	a := make(chan int, 50)
	b := make(chan int, 100)
	wg.Add(2)
	go f1(a)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}
