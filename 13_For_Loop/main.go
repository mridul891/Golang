package main

import "fmt"

// There is no other loop except For loop in GOlang

func main() {
	fmt.Println("For loop in Golang")

	// Basic for Loop
	for i := 0; i < 10; i++ {
		fmt.Println("Number is :", i)
	}

	// Infinte for loop with Break statement
	counter := 0
	for {
		fmt.Println("Infinite Loop")
		counter++
		if counter == 3 {
			break
		}
	}

	// In for loop we can use break to exit a loop prematurely or continue to move to the next iteration

	// The range Keyword simplifies looping  over elements of a collection like slices , arrays ,maps and strings.

	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Printf("Index of Number is:%d ,and value is %d\t", index, value)
	}

	fmt.Printf("\n")

	data := "Hello!"
	for index, char := range data {
		fmt.Printf("Index of Number is:%d ,and char is %c \t", index, char)
	}

	// range keyword in go simplifies the process of iterating over elements in various types of collections, such as slices ,array , maps and strings .It provides both the index (or key in case of map )and the values of each element in the collection .

	// In this example , range numbers return both the index and value of each element in the numbers slice. The loop will itereate five times (the length of the slice ) and on each iteration the index and value will be set to the current index and value respectively.
}
