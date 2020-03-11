package main

import (
	"fmt"
	"strings"
)

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() int {
	fmt.Println(distribution)
	sum := 0
	for _, value := range users {
		for _, v := range strings.Split(value, "") {
			switch {
			case v == "e" || v == "E":
				sum++
			case v == "i" || v == "I":
				sum += 2
			case v == "o" || v == "O":
				sum += 3
			case v == "u" || v == "U":
				sum += 4
			}
		}

	}
	return 50 - sum

	//sum := 0
	//for _, value := range users {
	//	for _, v := range strings.Split(value, "") {
	//		switch {
	//		case v == "e" || v == "E":
	//			sum++
	//		case v == "i" || v == "I":
	//			sum += 2
	//		case v == "o" || v == "O":
	//			sum += 3
	//		case v == "u" || v == "U":
	//			sum += 4
	//		}
	//	}
	//}
	//return 50 - sum
}

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}
