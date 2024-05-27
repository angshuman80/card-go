package main

import "fmt"

type card struct {
	card  string
	color string
}

type myDeck []card

func createStructDeck() myDeck {

	newStructDeck := myDeck{}

	cardDeckName := []string{"Diamond", "Spade", "Club", "Heart"}
	cardDecColor := []string{"Black", "red"}

	for _, name := range cardDeckName {
		for _, color := range cardDecColor {
			mycard := card{card: name, color: color}
			newStructDeck = append(newStructDeck, mycard)
		}
	}

	return newStructDeck
}

func (d myDeck) printDeck() {
	for _,name := range d {
		fmt.Println(name.card ," ",name.color)
	}
}