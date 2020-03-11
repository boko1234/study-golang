package main

import "fmt"

func assign(a interface{})  {
	fmt.Println(a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("字符串",str)
	}
}
func assign2(a interface{}) {
	fmt.Println(a)
	switch t := a.(type) {
	case string:
		fmt.Println("是字符串",t)
	case int:
		fmt.Println("是int",t)
	case int64:
		fmt.Println("int64",t)
	}
}
func main() {
	assign2(true)
	assign2("aaaaa")
	assign2(int64(200))

}
