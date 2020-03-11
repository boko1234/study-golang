package main

import "fmt"

type persona struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	var p persona
	p.name = "菜虚鲲"
	p.age = 18
	p.gender = "女"
	p.hobby = []string{"唱", "跳", "rap", "篮球"}
	fmt.Println(p)

	var s struct{
		c,x string
		y int
	}
	s.c = "哈哈哈"
	s.x = "嘿嘿嘿"
	s.y = 100
	fmt.Printf("type:%T value:%v\n",s,s)
	var ds = new(person)
	ds.name = "菜虚鲲"
	fmt.Printf("%T\n",ds)
	fmt.Printf("%p\n",ds)
}
