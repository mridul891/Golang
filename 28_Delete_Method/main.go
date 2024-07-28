package main

import (
	"fmt"
	"net/http"
)

func PerformDeleteMethod(myurl string) {
	req, err := http.NewRequest(http.MethodDelete, myurl, nil)
	if err != nil {
		fmt.Println("Error while Deleting ", err)
		return
	}

	//	Send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error while Deleting ", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("status", res.Status)
}

func main() {
	fmt.Println("learning Delete method in Golang")
	myurl := "https://jsonplaceholder.typicode.com/todos/2"
	PerformDeleteMethod(myurl)
}
