package main

import "fmt"

// a slice is a flexible and dynamic data structure that provides a more powerful alternative to arrays
// slice have a dynamic size and their length can be changed during the program's execution

func main() {
	fmt.Println("Slices in golang")
	// Declare and initialize a slice

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice:", numbers)

	// Determination of the data type
	fmt.Printf("Slice has data type : %T\n", numbers)

	fmt.Println("Length of the Slice", len(numbers))

	// append in slice

	numbers = append(numbers, 3, 10, 12, 34)
	fmt.Println("Slice:", numbers)

	// Intiatilization of empty slice
	name := []string{}
	fmt.Println("name:", name)
}
