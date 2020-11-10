package main

import "fmt"

// ContactInfo for a person
type contact struct {
	email   string
	phone   string
	zipCode string
}

// Person is a person class
type Person struct {
	firstName string
	lastName  string
	contact
}

func (p Person) print() {
	fmt.Printf("%+v", p)
}

func (p *Person) updateName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	tim := Person{firstName: "Tim", lastName: "U", contact: contact{
		email:   "jim@jim.com",
		phone:   "404-404-4400",
		zipCode: "99900",
	}}
	tim.print()
	tim.updateName("Jimmy")
	tim.print()
	// tim.lastName = "Anderson"
	// tim.print()
}
