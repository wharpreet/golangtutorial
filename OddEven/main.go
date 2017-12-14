package main

import (
	"fmt"
)

func main() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, i := range num {
		if i%2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
	}
}
