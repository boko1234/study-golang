package main

import "fmt"

type myInt int

func (m myInt) hello() {
	fmt.Println(m,"我是一个int")
}

func main() {
	m := myInt(100)
	fmt.Println(myInt(30))
	//m := 10
	m.hello()
}
