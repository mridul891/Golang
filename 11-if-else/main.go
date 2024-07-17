package main

import "fmt"

func main() {
	x := 10
	if x > 5 {
		fmt.Println("X is greater than 5")
	} else {
		fmt.Println(" X  is less than 5")
	}

	z := 10
	if z > 10 {
		fmt.Println("z is Greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than 10")
	} else {
		fmt.Println("z is less than 5")
	}
}
