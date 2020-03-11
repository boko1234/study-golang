package intee

import "fmt"

type animal interface {//洗衣机
	mover
	eater
}


type mover interface {//洗衣服功能
	move()
}

type eater interface {//烘干功能
	eat()
}

type cat struct {
	name string
	feet int8
}
//cat实现了mover接口
func (c *cat) move() {
	fmt.Println("猫")
}
//cat实现了eater接口
func (c *cat) eat(food string) {
	fmt.Println("猫吃",food)
}

func main() {

}





