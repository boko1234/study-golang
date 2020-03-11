package main

import "fmt"

// dog 这是一个狗的结构体
type dog struct {
	name string
	age int
}

func newDog(name string,age int) dog {
	return dog{
		name: name,
		age:age,
	}
}

func (d *dog) wang() {
	fmt.Println(d.name, ":汪汪汪~")
	d.age++
}

func main() {
	d1 := newDog("zhoulin",10)
	d1.wang()
	d1.wang()
	fmt.Println(d1.age)
}
