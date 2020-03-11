package main

import "fmt"

type car interface {//接口，把车放到对应的赛道
	run()
}

type falali struct {//造车
	brand string

}

type baoshijie struct {//造车
	brand string
}

func (f falali) run() {//法拉利专属
	fmt.Println(f.brand)
}

func (b baoshijie) run() {//保时捷专属
	fmt.Println(b.brand)
}




func drive(c car)  {//让车跑
	c.run()
}

func main() {
	var	f1 = falali{
		brand: "法拉利",
	}
	var b1 = baoshijie{
		brand: "保时捷",
	}
	drive(f1)
	drive(b1)
}
