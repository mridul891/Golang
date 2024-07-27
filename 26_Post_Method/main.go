package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func PerformPostMethod(myurl string) {
	todo := Todo{
		UserID:    23,
		Id:        201,
		Title:     "Hello Moto",
		Completed: true,
	}

	jsondata, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling :", err)
		return
	}
	// convert json data to string
	jsonString := string(jsondata)

	// convert string data to reader
	jsonReader := strings.NewReader(jsonString)

	// send Post Request
	res, err := http.Post(myurl, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error while Sending Data :", err)
		return
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while Sending Data :", err)
		return
	}

	fmt.Println("Response :", string(data))
}

func main() {
	myurl := "https://jsonplaceholder.typicode.com/todos"
	PerformPostMethod(myurl)
}
