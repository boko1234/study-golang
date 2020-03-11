package main

import "fmt"
//方法
type person struct {
	name string
	age  int
}

//type dog struct {
//	name string
//}
//构造函数
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func newDog(name string) dog {
	return dog{
		name: name,
	}
}

func main() {
	p1 := newPerson("斯大林", 10)
	p2 := newPerson("元首", 23)
	p3 := newDog("汪汪")
	fmt.Println(p1, p2, p3)
}

