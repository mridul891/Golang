package main

import "fmt"

func main() {

	// one way
	var name string = "golang"
	fmt.Println(name)

	// another way
	//  Here golang  by default  infer its  type
	var isAdult = true
	var name2 = "Mridul"
	var num int = 1
	fmt.Println(name2)
	fmt.Println(isAdult)
	fmt.Println(num)

	// short  hand syntax
	name3 := "Golang"
	fmt.Println(name3)

	// if we are starting the name in small letter then we are making the variable as private variable and it is usable in that package only
	name4 := "This is an private variable"

	// if we are starting the name in capital letter then we are making the variable as Public variable and it is usable in all the packages
	Name := "This is an Public variable"

	fmt.Println(name4)
	fmt.Println(Name)

}
