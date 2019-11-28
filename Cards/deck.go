package main

import "fmt"

//Create a new type of 'deck'
//Which is a slice of string

type deck []string

func newDeck() deck { //it will return a value of type 'deck'
	cards := deck{}

	cardSuits := []string{"Spades", "Daimonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
