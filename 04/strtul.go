package main

import "fmt"

type address struct {
	province string
	city     string
}

type peson struct {
	name string
	age  int
	address
}
type company struct {
	name string
	address
}
type ddd struct {
	int
	string
}

func main() {
	p1 := peson{
		name: "周林",
		age:  20,
		address: address{//写一样会报错
			province: "山东",
			city:     "威海",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.city)//这样写可以直接访问嵌套下面的数据
	fmt.Println(p1.address.city)//这样写可以直接访问嵌套下面的数据
	p2 := ddd{
		int:    10,
		string: "30",
	}
	fmt.Println(p2)
}
