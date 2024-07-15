package main

import "fmt"

// no input parameters , types and output data types
func simpleFunction() {
	fmt.Println("Simple function with no input and output parameters")
}

// with input parameter
// func functionname ( input parameters then input parameters type ) (output /return type){
// body of fuctions
// }
func add(a, b int) int {
	return a + b
}

// return the  named value from the function
func addwithname(a, b int) (result int) {
	result = a + b
	return result
}

func main() {

	fmt.Println("we are learning function in Golang")
	simpleFunction()
	ans := add(3, 4)

	fmt.Println("The ans is ", ans)
	ans1 := addwithname(4, 5)

	fmt.Println("The ans is ", ans1)

}
