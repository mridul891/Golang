package main

import "fmt"

// array is nothing but a collection of data of similar data type

// Array Declaration

// array are declared using the syntax var name [size]type , where name is the array variable , size is the number of elements and type is the data type of elements.

// arraus have a fixed size meaning you must specify te number of elements the array can hold at the time of declaration.

// once array is created , its size cannot be changed.

// example var number [5] int declares and integer array with 5 elements as an size.

// Array Intialization

// var number = [5] int  {1,2,3,4,5}

func main() {
	fmt.Println("Learning array in Golang")

	var name [5]string

	name[0] = "aakash"
	name[1] = "Mridul"
	name[4] = "Pandey"

	fmt.Println("Names of person is :", name)

	var number = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Intialized array  :", number)

	fmt.Println("Length of array is ", len(number))

	//  to access value at given index
	fmt.Println("vale of name at 2 index is", number[2])

	// In go , when you declare an array or a slice , the elements are intialized to their zero values , The zero values is the default values assigned to variables of a certain type when no explicit values is provided

	// In go if we just declared the array then it is initialized by with the default values of the data type
	// if int = 0
	// bool = false
	// string =""

	fmt.Printf("name  array is %q\n :", name)
}
