package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDeck(t *testing.T) {

	card := createDeck()
	fmt.Println(len(card))
	if len(card) != 16 {
		t.Errorf(" Was expecting 16 card but got %v",len(card))
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("test_file")
	card := createDeck()
	success := saveToFile(card,"test_file")
	assert.Equal(t,true,success)
    if !success{
		t.Errorf("Unable to save file")
	}
    os.Remove("test_file")
}

func TestReadfromFile(t *testing.T){
	os.Remove("test_file")
	card := createDeck()
	success := saveToFile(card,"test_file")
	assert.Equal(t,true,success)
    cardRes,err := readFromFile("test_file")
	assert.Equal(t,nil,err)
	cardResArr := strings.Split(cardRes, ",")
	assert.Equal(t,16,len(cardResArr))
	os.Remove("test_file")
}