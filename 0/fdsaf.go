package main


import "fmt"

func main()  {
	age := 10
	if age > 18 {
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
	if age > 40 {
		fmt.Println("中年")
	}else if age > 50 {
		fmt.Println("退休")
	}
}