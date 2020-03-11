package main

import (
	"fmt"
	"time"
)

func main() {
	for i :=0; i < 100; i++{
	    //go func(i int) {
	    //	fmt.Println(i)
		//}(i)
	    //go hello(i)
	    //go func() {
	    //	fmt.Println(i)
		//}()
	}
	 //hello()
	fmt.Println("main")
	time.Sleep(time.Second)
}

func hello(i int)  {
	fmt.Println("hello",i)
}