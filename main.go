package main

import "fmt"

func main() {
   //peronStruct()
   //mapConstruct()
   //shapeInterface()
   //readJsonTest()
   callChannel()
}

func readJsonTest(){

   readJson("https://jsonplaceholder.typicode.com/todos/1")
	reaSampleJson()
	readSimpleJson()
	//readJson()
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

func mapConstruct(){
  myMap :=	creatMap()
  printMap(myMap)
}

func shapeInterface(){
	// Get area of square
	sq := square{sideLength: 10}
	fmt.Println("Square area is ")
	printArea(sq)

	// Get area of Triangle
	tr := triangle{base: 10,height: 20.5}
	fmt.Println("Triangle area is ")
	printArea(tr)
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
