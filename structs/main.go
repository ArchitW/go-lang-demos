package main

import "fmt"
//go is pass-by value.
//more of less like object in JS or Dicts in Python
/*

value types: int, float, string,bool, struct
reference types: slice, map, channel, pointer, functions

 */
type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contactInfo // adding struct in struct, can simply put contactInfo
}

func main() {

	archit := person{
		firstName: "Archit",
		lastName:  "W",
		contactInfo: contactInfo{
			email:   "mail@mail.com",
			zipCode: 44434,
		},
	}

	fmt.Printf("%+v", archit)

	/* we can not update name using this method
	because when we call like this, it will create reference to current variable and assign new value to it
	archit.updateName("__Archit__")
	archit.print()
	func (p person) updateName(newFname string)  {
		p.firstName = newFname
	}
	*/
	/* update struct value with pointers */
	/*
		pointerVar := &archit
		pointerVar.updateName("__Updated Name__")
	is same as
	archit.updateName("___Update Name__)
	*/

	archit.updateName("New Name")
	archit.print()
}

func (p person) print() {
	fmt.Println(p)
}

func (pointerToPerson *person) updateName(fName string) {
	(*pointerToPerson).firstName = fName
}

/*


//Create and Access Person struct
john := person{"john","doe",12} // careful with this type of declaration, prune to confusion or manipulation later
//or
smith := person{firstName:"smith", lastName:"cena"}
//og
var alex person
fmt.Println(alex)
fmt.Printf("%+v", alex) // print field names and values
fmt.Println("---x---")
//accessing

alex.firstName = "test"
alex.lastName = "PPP"


fmt.Println(john)
fmt.Println(smith)


Pointers:

address to value => *address
value to address => &value

Receivers:
(t *type) looking for a type pointer


*/
