package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func (p person) print() {
	fmt.Printf("\n%+v\n", p)
}

// Buradaki * tipi belirtiyor
// Normalde başında yıldız varsa pointer'ın value'suna denk gelir demek
// & işareti varsa value'nun pointer'ına denk gelir demek
func (p *person) updateNamePointer(firstName string) {
	// p.firstName = firstName
	(*p).firstName = firstName
}

func main() {
	var person1 person
	person1.print()

	person0 := person{
		"Alex",
		"Anderson",
		contactInfo{email: "-", zipCode: 0},
		contactInfo{email: "-", zipCode: 0},
	}
	person0 = person{firstName: "Alex", lastName: "Anderson"}
	person0.firstName = "John"

	// person0Pointer := &person0
	// person0Pointer.updateNamePointer("Hans")
	person0.updateNamePointer("Hans")
	person0.contact = contactInfo{email: "email", zipCode: 35}

	person0.print()
}
