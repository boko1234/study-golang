package main

import "fmt"

//func adder() func(int) int {
//	sum := 0
//	return func(x int) int {
//		sum += x
//		return sum
//	}
//}
//
//func main() {
//	pos, neg := adder(), adder()
//	for i := 0; i < 10; i++ {
//		fmt.Println(pos(i), neg(-2*i))
//	}
//}



//TODO:

func adder() func(int) int {
	x := 100
	return func(i int) int {
		fmt.Println(x)
		x += i
		fmt.Println(x)
		return x
	}
}





func main() {
	ret := adder()
	//ret3 := ret(200)
	fmt.Println(ret(100))
	//fmt.Println(ret3)
}