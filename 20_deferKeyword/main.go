package main

import "fmt"

func main() {
	fmt.Println("Defer keyword in Golang")

	//  Defer  makes the function to display the output just before when its parent function fineshes its execution

	fmt.Println("Starting of the program")
	defer fmt.Println("Middle of the program")
	defer fmt.Println("End of the program")

}

// If 2 defer keyword is used in a function then it goes in the lifo state or inotherwise when two defer keyword are used it store the function calling in a stack .
