package main

import "fmt"

// function mai hamesha ek pointer ayega jab bhi address ayega as an input from the main function
func modifyvalueByRefrence(num *int) {
	*num = *num + 30
}

func main() {
	fmt.Println("Pointers in Golang ")

	// Pointers are nothing just a varibale which stores the address of the other variable
	// Pointers are used to indirectly refer to the value stored in variable rather than its value itself
	// To declare a pointer we user the * symbol followed by the type of the variable it will point to . you can then intialize the pointer using the address of (&) operator

	// Declare a variable and a pointer

	num := 3

	ptr := &num

	fmt.Println("Num has value : ", num)
	fmt.Println("ptr has value : ", ptr)
	fmt.Println("Data in ptr has value : ", *ptr)

	// In Go pointers are intialized with nil by default if not explicity set to point to a valid memory address. A nil pointer doesnot point to any valid memory location

	var Pointer1 *int
	if Pointer1 == nil {
		fmt.Println("Pointer is nil / empty in nature")
	}

	// passing of pointer to a function

	value := 4
	modifyvalueByRefrence(&value)
	fmt.Println("Value contains :", value)
}
