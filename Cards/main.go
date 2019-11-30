package main

import "fmt"

func main() {
	deck := newDeckFromFile("my_cards")
	fmt.Println(deck)
}
