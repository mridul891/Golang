package main

import "fmt"

// In Go, the switch statement provides a way to conditionally execute code based on
// the value of an expression. It's a more concise and flexible alternative to a sequence of
// if-else statements when dealing with multiple possible values for a variable.

func main() {

	fmt.Println("Switch Case in Golang")

	day := 3

	// switch with single values
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	default:
		fmt.Println("unKnown day")
	}

	// switch with multiple values
	months := "Jan"

	switch months {
	case "Jan", "Feb", "mar":
		fmt.Println("Winter")
	case "April", "may", "june":
		fmt.Println("spring")
	default:
		fmt.Println("unKnown season")
	}

	// switch with expression

	temperature := 25

	switch {
	case temperature < 0:
		fmt.Println("Winter")
	case temperature >= 0 && temperature < 10:
		fmt.Println("cold")
	default:
		fmt.Println("unKnown season")
	}

}
