package main

import "fmt"

func Factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func taijie(n uint64) uint64 {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return taijie(n-1) + taijie(n-2)
}

func main() {
	//var i int = 6
	//fmt.Printf("%d的阶乘是 %d\n", i, Factorial(uint64(i)))
	ret := taijie(6)
	fmt.Println(ret)
}
