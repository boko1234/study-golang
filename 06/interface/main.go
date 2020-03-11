package main

import "fmt"

type speaker interface {
	speak()//只要实现了speak方法的变量都是speaker类型
	//方法签名
}

type cat struct {
}

type dog struct {
}
type person struct {
}

func (d dog) speak() {
	fmt.Println("汪汪汪")
}

func (c cat) speak() {
	fmt.Println("喵喵喵")
}
func (p person) speak() {
	fmt.Println("嘤嘤嘤")
}
func da(x speaker) {
	x.speak()

}

func main() {
	var c1 cat
	var d1 dog
	var p1 person
	da(d1)
	da(c1)
	da(p1)


	var ss speaker
	ss = c1
	da(ss)
	fmt.Println(ss)

}
