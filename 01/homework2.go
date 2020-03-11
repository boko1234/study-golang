package main

import "fmt"

func main() {
	for i :=0; i < 10; i++{
	    i = 3
	}


	ss := "一二三四五四三二一"
	fmt.Println([]rune(ss))
	fmt.Println(make([]rune, 0, len(ss)))
	r := make([]rune, 0, len(ss))
	for _, c := range ss {
		r = append(r,c)
	}
	fmt.Println("rune" ,r)
	for i := 0; i < len(r)/2 ; i++{
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}

	}
	fmt.Println("是回文")
}
