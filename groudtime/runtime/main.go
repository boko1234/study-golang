package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func a()  {
	defer wg.Done()
	for i :=1; i < 10; i++{
	    fmt.Println("A",i)
	}
}

func b() {
	defer wg.Done()
	for i :=1; i < 10; i++{
	    fmt.Println("B",i)
	}
}

func main() {
	maxProbes := runtime.NumCPU()
	if maxProbes > 1 {
		maxProbes -= 1
	}
	fmt.Println(maxProbes)
	runtime.GOMAXPROCS(maxProbes)
	fmt.Println(runtime.NumCPU())
	wg.Add(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
