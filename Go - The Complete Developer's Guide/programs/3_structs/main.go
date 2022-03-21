package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Jefferson",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}

	jim.updateName("Jimmy")
	jim.print()
}

func (personPointer *person) updateName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
