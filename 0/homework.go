
package _

import "fmt"

func main()  {
	for i :=0; i < 10; i++{
	    fmt.Print(i)
	}
	s := "hello"
	for v := range s  {
		fmt.Print(v)
		fmt.Print(s)

	}
	for i :=1; i < 10; i++{
		fmt.Printf("%d\t",i)
	}



}
