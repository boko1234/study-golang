package main

import "fmt"

const a = "hallowed"
const (
	n1 = iota
	n2
	_
	n4
)
const (
	d1,d2 = iota+1,iota+2
	d3,d4 = iota+1,iota+2
)

const (
	kb = 1<<(10*iota)
	md = 1<<(10*iota)
	yb = 1<<(10*iota)
	pb = 1<<(10*iota)
)

const (
	aa = 100
	b = 10324
)




func main() {
	b1 := true
	fmt.Printf("%T\n",b1)
}
