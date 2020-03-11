package main

import "fmt"

func main() {
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 16)
	m1["name"] = "周林"
	m1["age"] = 40
	m1["merried"] = true
	m1["hobby"] = [...]string{"唱"}
	fmt.Println(m1)
}
