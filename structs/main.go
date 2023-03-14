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
	//alex := person{"Alex", "Anderson"} //Because firstName is first in struct definition,
	//"Alex" is assigned to firstName. THIS IS FRAGILE
	//alex := person{firstName: "Alex", lastName: "Anderson"} //Explicit definition is more robust
	/*
		var alex person //Defines alex as type person with zero values assigned. "" in case of string
		alex.firstName = "Alex"
		alex.lastName = "Anderson"
		fmt.Println(alex)
		fmt.Printf("%+v", alex)
	*/
	jim := person{
		firstName: "jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // Even last line must have a comma...
		}, // ...Yes, even *here*
	}
	//fmt.Printf("%+v", jim)
	//	jimPointer := &jim // Create variable containing pointer to jim struct
	// & is an operator to give access to memory address var is pointing at
	// ie "referencing" of pointer
	//	jimPointer.updateName("jimmy") // Pass pointer to updateName function
	// Instead of this ^^^ and instantiating jimPointer with &jim
	// This instead vvv and no other change needs to be made in the printPer func
	jim.updateName("jimmy") // Pass pointer to updateName function
	jim.printPer()
}

func (p person) printPer() {
	fmt.Printf("%+v", p):w http.ResponseWriter, r *http.Request
}

// Go is a pass by value language. Whenever a value is passed to a function, that value is copied
// in whole to a different memory address. A new pointer is created.
// Doesn't apply to slices. Slice data structure has a ptr to the head of the array containing the
// slice elements. Go still makes a copy of that slice data, but that has a copy of the head of the array
// containing the slice elements. It is a "reference" type.
// Reference types: slices, maps, channels, pointers, functions
// Pointer types: int, float, string, bool, structs

func (pointerToPerson *person) updateName(newFirstName string) { // *person is a type description here,
	// not a dereferenced pointer. It is a pointer type, specific to a pointer to a person struct type...
	// ...OR if the variable is type person rather than *person, Go just handles it.
	(*pointerToPerson).firstName = newFirstName // *pointerToPerson gives value of variable "jim"
	// * is here an operator that gives the value of the variable referenced by the pointer
	// ie "dereferencing" the pointer
}

/*
Turn a value into an address with &value
Turn an address into a value with *address
*/
