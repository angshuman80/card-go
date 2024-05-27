package main

import "fmt"

func main() {
   peronStruct()
}

func deckWithoutStruct(){
	card := createDeck()
	card.print()
	fmt.Println(" Read and Write from file")
    success := saveToFile(card,"my_file")

	if success {
		result,_ := readFromFile("my_file")
		fmt.Printf("Response from file %v",result)
	}
}

func deckWithStruct(){
	cardStruct := createStructDeck()
	cardStruct.printDeck()
}

func peronStruct(){
	p := populatePerson()
	addressP := &p
	fmt.Println("addess ",addressP)
	fmt.Println(p)
	addressP.updatePerson("SarojK")
	fmt.Println("After update with pointer",p)
	p.updatePerson("alex")
	fmt.Println("After update ",p)
	pUpdated := updatePersonWihoutPointer("Saroj",p)
	fmt.Println("After update without Pointer ",pUpdated)

}
