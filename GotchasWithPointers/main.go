package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "How", "Are", "You"}
	updateSlice(mySlice)
	// This will replace the value in the slice
	// Quite opposite to how it works in strucks
	// Coz slice is a reference data type
	fmt.Println(mySlice)
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
