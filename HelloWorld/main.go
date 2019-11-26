package main

//package is collection of common source code file
//This file belongs to package called "main" making it an executable package.
//If we give some other package name, it will become a "resuable package"
//An executable package must have a function "main"
//If name is not main, on running file, we get error - 'go run: cannot run non-main package'

import "fmt"

//We are importing "fmt" package. fmt stands for "format"

func main() {
	fmt.Println("Hello World")
}
