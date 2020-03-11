package main

import "fmt"

//类型断言

func main() {
	var a interface{}
	a = 100
	//如何判断a保存的值
	//类型断言

	v, ok := a.(int8)
	if !ok {
		fmt.Println("猜对了", v)
		return
	} else {
		fmt.Println("猜错了")
	}
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8",v2)
	case int16:
		fmt.Println("int16")
	case string:
		fmt.Println("strign")
	case int:
		fmt.Println("int")
	default:
		fmt.Println("没了")
	}

}
