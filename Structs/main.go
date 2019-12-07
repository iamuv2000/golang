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
	yuvraj := person{
		firstName: "Yuvraj",
		lastName:  "Singh",
		contact: contactInfo{
			email:   "yuvi@gmail.com",
			zipCode: 6500013,
		},
	}
	//yuvrajPointer := &yuvraj //'&' Give me the address of variable "yuvraj"
	yuvraj.updateFirstName("Yuvi")
	yuvraj.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName //'*' Give me the value of the thing existing at this memory address
	// person is a 'type' of data
	//hence, in front of person when we use '*'
	//we refer to a new data type- a pointer that points to the data of that data type.
	//ie, *person => is a type of data, pointing to the address of type person
}

func (pointerToPerson *person) print() {
	fmt.Println(*pointerToPerson)
}
