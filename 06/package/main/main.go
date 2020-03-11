package main

import (
	"../cale"
	"fmt"
)
//如果包中的标识符小写，表示私用，
//大写公用
func main() {
	ret := calc.Add(100,30)
	fmt.Println(ret)
}
