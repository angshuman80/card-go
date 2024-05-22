package main

import "fmt"

func main() {
	card := createDeck()
	card.print()
	fmt.Println(" Read and Write from file")
    success := saveToFile(card,"my_file")

	if success {
		result,_ := readFromFile("my_file")
		fmt.Printf("Response from file %v",result)
	}
}
