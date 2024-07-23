package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Data conversion in Golang")

	// Data conversion is process of converting a value from one data type to another

	// Number Type Conversion
	num := 2
	fmt.Printf("Type of num is %T\n", num)
	var data float64 = float64(num)
	fmt.Println("Data is ", data)
	fmt.Printf("Type of Data is %T\n", data)

	// string Conversion
	// Itoa = integer to alphabet

	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is ", str)
	fmt.Printf("Type of Data is %T\n", str)

	// interger Convertor
	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("number_int is ", number_int)
	fmt.Printf("number_int of Data is %T\n", number_int)

	num_string := "3.14"
	num_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("num_float is ", num_float)
	fmt.Printf("num_float of Data is %T\n", num_float)

}
