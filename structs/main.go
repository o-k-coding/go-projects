package main

import "fmt"

// Properties are defaulted to values according to the type
// For instance string is assigned to ""
// numbers are assigned to ""
// bools are assigned to false

// This is important, the values are not undefined by default (or set to nil)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// you can also do this
	// and it creates the field name of contactInfo with type of contactInfo
	// contactInfo
}

func main() {
	// Assumes that the field ordering matches the value ordering
	// ALso all properties are required to be filled out for this method
	mike := person{"Mike", "O'Keefe", contactInfo{"my-email@email.com", 12345}}
	// This also works fine
	kass := person{firstName: "Kass", lastName: "O'Keefe"}

	// This will create a new instance of the struct with the default values
	var ollie person

	// & gives the memory address of the value this variable is pointing at (get the reference/pointer)
	mikePointer := &mike

	mikePointer.updateName("Mikey")

	// can also do this, Go automatically passes the pointer if the receiver function signature
	// accepts a pointer
	mike.updateName("MIKEY")

	mike.print()
	kass.print()
	ollie.print()
}

// By default Go passes by value
// so each time this is called, Go copies the person struct
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// Passing the pointer to a person ensures the update happens on the actual struct
// So the pointer will be the reference or address for the variable
func (pointerToPerson *person) updateName(newFirstName string) {
	// Using * on the pointer gets the value from the address of the pointer
	(*pointerToPerson).firstName = newFirstName
}

// So basically by default Go is going to be "immutable", but by using pointers you can make it mutable
// Go has the concept of value vs reference types
// For instance a struct is a value type and slices are a reference type
// With slices, they store the length, capacity and a pointer to an array
// So when passed to a function, the slice is copied, but the array is not because it is referenced, not stored by the slice
// This makes the slice a reference type

func (p person) print() {
	fmt.Printf("%+v", p)
}
