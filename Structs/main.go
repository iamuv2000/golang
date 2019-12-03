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
	fmt.Println(yuvraj)
}
