package main

import "fmt"

func main() {
	cards := []string{"Ace of Daimonds", newCard()}
	cards = append(cards, "Six of spades")
	for card, i := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Daimonds"
}
