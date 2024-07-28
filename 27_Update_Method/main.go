package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	Id        int    `json:"id"`
	UserId    int    `json:"userId"`
	Title     string `json:"Title"`
	Completed bool   `json:"completed"`
}

func performUpdateRequest(myurl string) {
	todo := Todo{
		UserId:    23,
		Title:     "Mridul Pandey JOd",
		Completed: false,
	}

	//	converting this data to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshalling : ", err)
		return
	}

	//convert json to string
	jsonstring := string(jsonData)

	//Converting this jsonstring to a reader because newrequest requires a reader

	jsonReader := strings.NewReader(jsonstring)

	//	Crete Put request
	req, err := http.NewRequest(http.MethodPut, myurl, jsonReader)
	if err != nil {
		fmt.Println("Error while Updating the details", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	//	Send The request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while sending request", err)
		return
	}
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println("Response :", string(data))
	fmt.Println("Status", res.Status)
}
func main() {
	fmt.Println("learning Update Request in Golang")
	myurl := "https://jsonplaceholder.typicode.com/todos/2"
	performUpdateRequest(myurl)
}
