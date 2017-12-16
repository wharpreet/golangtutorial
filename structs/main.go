package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	//	alex := person{firstName: "Alex", lastName: "Anderson"}
	//	fmt.Println(alex)

	//	var alex person
	//	alex.firstName = "Alex"
	//	alex.lastName = "Anderson"
	//	fmt.Println(alex)
	//	fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//jimPointer := &jim
	//jimPointer.updateName("Jimmy")

	jim.updateName("Jimmy")

	//	jim.updateName("Jimmy")
	jim.print()
}

func (pointerToPersion *person) updateName(newFirstName string) {
	(*pointerToPersion).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)
}
