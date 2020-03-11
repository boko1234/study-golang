package main

import "fmt"

type person struct {
	name string
	age  int
}


type addr struct {
	name string
	age  int
}

//指针，不使用的话只会修改副本
func (p *addr) dream() {
	p.age++
	fmt.Println(p.name, p.age)

}

func newPerson(n string, i int) person {
	return person{
		name: n,
		age:  i,
	}

}
func newAddr(n string, i int) addr {
	i++
	return addr{
		name: n,
		age:  i,
	}

}

//函数 谁可以用 函数的名称（参数）（返回值）{内容}
//func () fsdfd() (fff int){}
func main() {
	a := newAddr("卧槽", 100)
	a.name = "哈喽"
	fmt.Println(a)
	a.dream()
	fmt.Println(a)
}
