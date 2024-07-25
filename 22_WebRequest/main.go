package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// 	defer response. Body. Close ( ) statement is used to ensure that the response
	// body is closed after you have finished reading from it. In Go, it's important to close
	// resources like network connections and file handles to free up system resources.

	// Resource Management: HTTP responses in Go are represented by http. Response
	// objects, which have a Body field containing the response body. This body is a stream of
	// data from the server, and it's important to close this stream once you're done reading
	// from it to release the associated resources.

	// ioutil.ReadAll(response. Body): This line reads the entire response body
	// (response. Body) into a byte slice (body). ioutil. ReadAll reads from the
	// response. Body until an EOF (end of file) is reached and returns the data read and an
	// error. If there's an error reading the response body (e.g., connection closed, invalid
	// data), it will be stored in err.
	
	// fmt.Println(string (body)): This line converts the byte slice body to a string and
	// prints it to the console. The response body typically contains the content (data) returned
	// by the server in response to the request.

	fmt.Println("Learning Web Request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2")

	if err != nil {
		fmt.Println("Error Getting Get response", err)
		return
	}

	defer res.Body.Close()
	fmt.Printf("Type of response is :%T\n ", res)
	// Gives  a  http response as response
	// fmt.Println("response", res)

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error Getting while Reading Data", err)
		return
	}
	// read all gives data the response in the bytes format we need to convert this into string
	fmt.Println("response:", string(data))

}
