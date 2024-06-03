package main

import (
	//"encoding/json"
	"bytes"
	"encoding/json"
	"fmt"

	//"io"
	//"log"
	"net/http"
)

type Message struct{
	UserId int
	Title string
	Id int
	Completed bool
}

type Operation struct{
	Operation string  `json:"operation"`
	Key string  `json:"key"`
}

type Squad struct {
	SquadName string `json:"squadName"`
	Formed    int    `json:"formed"`
	Leader    string `json:"leader"`
	Status    bool   `json:"active"`
	Members   []struct {
		Name           string `json:"name"`
		Age            int    `json:"age"`
		SecretIdentity string `json:"secretIdentity"`
	} `json:"members"`
}

// main function
func reaSampleJson() {

	// defining a struct instance
	var squad Squad

	// string json
	jsonString := `{
		"squadName": "Go Linux Cloud",
		"formed": 2022,
		"leader": "",
		"active": true,
		"members": [
		  {
			"name": "Anna",
			"age": 30,
			"secretIdentity": "Duke"
		  },
		  {
			"name": "Harry Potter",
			"age": 32,
			"secretIdentity": "Jane"
		  }]}`

	// decoding squad struct from json string
	err := json.Unmarshal([]byte(jsonString), &squad)

	if err != nil {
		// print out if error is not nil
		fmt.Println(err)
	}

	// printing details of struct
	fmt.Println("Struct is:", squad)
	fmt.Println("Squad's name is:", squad.SquadName)
	fmt.Println("Squad's leader is:", squad.Leader)
	fmt.Println("1st member is:", squad.Members[0].Name)
}

func readSimpleJson(){
	var data Operation
	jsonString := `{
		"operation": "Go Linux Cloud",
		"key": "imp"}`
    json.Unmarshal([]byte(jsonString), &data)
    fmt.Println(data)
}

func readJson(url string) {
	resp,err := http.Get(url)
	fmt.Println(resp.StatusCode)
	if err==nil && resp.StatusCode==200 {
		//fmt.Println("resp = ",resp)
		respByte := make([] byte,1024)
		resp.Body.Read(respByte)
	//	jsonResult := string(respByte)
	//	fmt.Println("Resp Body = ",string(jsonResult))
		var m Message
		err := json.Unmarshal(bytes.Trim(respByte,""),&m)
		if err!=nil{
			fmt.Println("Error while marshalling ",err)
		}
		fmt.Println(m.Title)

		//jsonResult := json.NewDecoder(resp.Body)
	/*	for {
			var m Message
			if err := jsonResult.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			fmt.Println("%s: %s\n", m.userId, m.title)
		}*/
	}
}