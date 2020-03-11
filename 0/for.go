package _

import "fmt"

func main() {

	//for   {
	//	fmt.Println("123")
	//}

	fmt.Println("hello")
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(i, "*", j, "=", i*j, " ")
		}
		fmt.Println()
	}
}
