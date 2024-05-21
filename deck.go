package main

import "fmt"

type deck []string

func createDeck() deck {
	newDeck := deck{}

	cardDeckName := []string{"of Diamond", "of Spade", "of club", "of heart"}
	cardDeckNum := []string{"One ", "Two ", "Three ", "Four "}

	for _, cNum := range cardDeckNum {
		for _, cName := range cardDeckName {
			name := cNum + cName
			newDeck = append(newDeck, name)
		}
	}

	return newDeck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}