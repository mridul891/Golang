package main

import (
	"encoding/json"
	"fmt"
)

// Ways to Write JSON

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_Adult"`
}

func main() {
	fmt.Println("JSON in Golang")
	person := Person{
		Name:    "Mridul ",
		Age:     20,
		IsAdult: true,
	}
	fmt.Println("Person Data is :", person)

	// conert person into Json Encoding (Marshalling)
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Error in Marshalling", err)
		return
	}
	fmt.Println("Person Data is :", string(jsonData))

	// decoding (unMarshalling)
	var PersonData Person

	err = json.Unmarshal(jsonData, &PersonData)
	if err != nil {
		fmt.Println("Error in UnMarshalling", err)
		return
	}
	fmt.Println("Person Data is  after unmarshalling" ,PersonData)
}
