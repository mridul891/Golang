package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning CRUD...")

	response, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error while Getting the data ")
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Println("Error in gettinf response", response.Status)
		return
	}

	// data, err := io.ReadAll(response.Body)
	// if err != nil {
	// 	fmt.Println("Error while Getting the data ")
	// 	return
	// }
	// fmt.Println("Data:", string(data))

	var todo Todo
	err = json.NewDecoder(response.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error while Getting the data ")
		return
	}
	fmt.Println("Todo :", todo)
}
