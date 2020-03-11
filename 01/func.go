package main

import "fmt"




//参数可以命名也可以不命名
//命名的返回值就相当于在函数中生命一个变量
func sum(x int, y int) (ret int) {
	return x + y
	//ret = x+y 可以省略return
}

func f2() {
	fmt.Println("f2")
}

func f3() int {
	return 3
}

func f7(x string, y ...int) {
	fmt.Println(x)
	fmt.Println(y)
}
func f6(x, y, z int, m, n string, i, j bool) {
	return
}
func fdf(a int,b ...string)  {
	fmt.Println(a)
	fmt.Println(b)
}


func main() {
	//r := sum(1,2)
	//fmt.Println(r)
	//f2()
	//fmt.Println(f3())
	fdf(3,"434","fds")
	f7("hello",1,2,3)
}




























