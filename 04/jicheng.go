package main

import "fmt"

type animal struct {
	name string
}

func (a animal) move() {

}

type dag struct {
	feet uint8
	animal
}

func (d dag) wang()  {
	fmt.Printf("%s在叫：汪汪汪~\n",d.name)
}

func main() {
	d1 := dag{
		feet:   4,
		animal: animal{name: "狗蛋"},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
