package main

import "fmt"

func main() {
	age := 21
	name := "mridul"
	height := 5.115

	fmt.Println("age:", age, "Height:", height, "name:", name)

	// printf is basically for formatter

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %f\n", height)

	// %T = Tell the type of the variable

	fmt.Printf("Type of name is %T\n", name)
	fmt.Printf("Type of height is %T\n", height)

	// for multiple formatter 
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
