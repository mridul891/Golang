package main

import (
	"fmt"
	"net/url"
)

func main() {

	fmt.Println("Learning Url ")
	myurl := "https://www.example.com/path/to/resource?key1=value1&key2=value2"
	fmt.Printf("Type of Url :%T\n", myurl)
	fmt.Println(myurl)

	// 	Parsing URL : The url.Parse function is used to parse a string into  Url Object
	parseUrl, err := url.Parse(myurl)
	if err != nil {
		fmt.Println(" Can't parse URL", err)
		return
	}
	fmt.Println(" URL is :", parseUrl)
	fmt.Printf("Type of the Url is : %T\n", parseUrl)

	// Accessing URL Components:
	//	Once we have the URL Object , we can access its components using various fields:
	// 	Scheme : Indicates the Protocol used (eg : "https")
	// 	Host : Specifies the domain name and optionally the port number
	//	Path : Represents the path component of the url which specifies the resources location on the server
	// 	RawQuery : Contains the raw query string, including query parameters.

	// Query Parameters :
	// Query parameters are the key value pairs appended to the end of a url usually starting with a ? and separated by & The url.values type is used to represent query parameters as a map.You can access and manipulate query parameters using methods like GET,SET and Add
	fmt.Println("Scheme :", parseUrl.Scheme)
	fmt.Println("Host :", parseUrl.Host)
	fmt.Println("Path :", parseUrl.Path)
	fmt.Println("Rawquery :", parseUrl.RawQuery)

	//	Modifying URL Components
	parseUrl.Path = "/newPath"
	parseUrl.RawQuery = "username=MridulPandey"

	//	Constructing a URl string from  url string
	newUrl := parseUrl.String()

	fmt.Println("newUrl is :", newUrl)

}
