package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	yuvraj := person{firstName: "Yuvraj", lastName: "Singh"}
	fmt.Println(yuvraj)
}
