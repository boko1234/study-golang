package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string
	Age int
}

func main() {
	p1 := person{
		Name: "狗蛋",
		Age:  30,
	}
	b,err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
	str := `{"name":"理想","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%v\n",p2)
}
