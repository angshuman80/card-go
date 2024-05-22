package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

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

func (d deck) toString() string{
	return strings.Join(d,",")
}

func saveToFile(d deck,fileName string) bool{
  err :=  os.WriteFile(fileName,[] byte (d.toString()),0666)

  if err!=nil{
	log.Fatal("Unable to write to file ",err)
	return false
  }
  return true
}

func readFromFile(fileName string) (string,error){
	deckArr , err :=  os.ReadFile(fileName)

	if err!=nil{
	  log.Fatal("Unable to read from file ",err)
	  return "",errors.New("Unable to read file")
	}
	  deckStr := string(deckArr)
	  return deckStr,nil
  }