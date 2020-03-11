package _

import "fmt"

func main() {
	var n = 3
	//if n == 1 {
	//	fmt.Println("1")
	//} else if n == 2 {
	//	fmt.Println("2")
	//} else if n == 3 {
	//	fmt.Println("3")
	//} else if n == 4 {
	//	fmt.Println("4")
	//} else if n == 5 {
	//	fmt.Println("5")
	//}
	switch n {
	case 1:
		fmt.Print("1")

	case 2:
		fmt.Print("2")

	case 3:
		fmt.Print("3")

	case 4:
		fmt.Print("4")

	case 5:
		fmt.Print("5")
	}
	for i := 0; i < n; i++ {
		fmt.Println(i)
		if i == 1 {
			break
		}
		fmt.Println(i)
	}
	fff := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, v := range fff {
		fmt.Print(",,",v)
		sum = sum + v
		fmt.Println(sum)
	}


	fmt.Print(fff[:4])

}
