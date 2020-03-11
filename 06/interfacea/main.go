package main

import "fmt"

type animal interface {
	move()
	eat(string)
}
type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("猫")
}

func (c cat) eat(st string) {
	fmt.Println("吃",st)
}

type chicken struct {
	feet int8
}

func (c chicken) move() {
	fmt.Println("鸡动")
}
func (c chicken) eat(st string) {
	fmt.Println("吃",st)
}

func main() {
	var a1 animal
	bc := cat{
		name: "淘气",
		feet: 3,
	}
	a1 = bc
	a1.eat("小黄鱼")
	fmt.Println(a1)
	kfc := chicken{feet:30}
	a1 = kfc
	a1.eat("面条")
	fmt.Println(a1)
}



