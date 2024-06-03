package main

import "fmt"

func creatMap() map[string]string {
	myMap := map[string]string{
		"name":    "Ang",
		"subject": "play",
		"age":     "40",
	}
	return myMap
}

func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("Key = ",key , " value ",value)
	}
}