package main

import "fmt"

// a slice is a flexible and dynamic data structure that provides a more powerful alternative to arrays
// slice have a dynamic size and their length can be changed during the program's execution

// use of make function : you can use the make function to create a slice with a specific length and capacity.
func main() {
	fmt.Println("Slices in golang")
	// Declare and initialize a slice

	numbers := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice:", numbers)

	// Determination of the data type
	fmt.Printf("Slice has data type : %T\n", numbers)

	fmt.Println("Length of the Slice", len(numbers))

	// Capacity of the Slice
	fmt.Println("capacity of the Slice", cap(numbers))

	// append in slice

	numbers = append(numbers, 3, 10, 12, 34)
	fmt.Println("Slice:", numbers)

	// Intiatilization of empty slice
	name := []string{}
	fmt.Println("name:", name)

	// Intialization of the slice using make function

	numberwithmake := make([]int, 3, 5)

	numberwithmake = append(numberwithmake, 10)
	numberwithmake = append(numberwithmake, 98)

	// If the capacity of the slice is full then it automatically double down its capacity

	fmt.Println(" Slice initization uising make function")
	fmt.Println("Slice", numberwithmake)
	fmt.Println("Length of the Slice", len(numberwithmake))
	fmt.Println("Capacity of the Slice", cap(numberwithmake))

}
