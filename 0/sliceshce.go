package _

import "fmt"

func main() {
	s1 := []string{"背景", "伤害", "深圳"}

	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s1 = append(s1, "广州")

	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))
}
