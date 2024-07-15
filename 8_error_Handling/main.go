package main

import "fmt"

// In GO we Normally handle Errors as we do in other Programming language, but again go Introduced a unique concept of _(underscore)

// In go , the (_) is used as a blank identifier , It serves as a write only variable that allows you to discard values returned by a function or to ignore specific values when you are not interested in using them

// Jaise error mil raha hai return but hume usse ignore karna hai to hum _ use kar sakte hai
func divide(a, b float64) (float64, error) {
	if b == 0 {
		// error should always start with small character
		return 0, fmt.Errorf("denominator must not be zero")
	}
	return a / b, nil
}

// Easy basha mai underscore tab use karna hai jab hum kissi function ek output ko ignore karna ho tab us case mai usse use kar sakte hai

// if we want to return any other data type

func dividereturnstring(a, b float64) (float64, string) {
	if b == 0 {
		// error should always start with small character
		return 0, "denominator must not be zero"
	}
	return a / b, "nil"
}

func main() {

	fmt.Println("started Error Handling in the function ")

	ans, _ := divide(10, 0)

	fmt.Println("Division of two number is ", ans)

	ans2, _ := dividereturnstring(10, 0)

	fmt.Println("Division of two number is ", ans2)
}

// // summary

// In this example:
// The divide function returns two values: the result of the division and an error (if any).

// In the main function, when calling divide (10, 2), we are only interested in the
// result, so we use the blank identifier (_) to discard the error.

// When calling divide (10, 0), we again use the blank identifier to discard the error, as
// we are not interested in handling it in this specific scenario.

// • Absolutely, you can use any valid variable name in place of the underscore (_). The use
// of the underscore is a convention in Go to indicate that the variable is intentionally being
// ignored, and it's a way to make it clear to both the compiler and other developers that
// you are not planning to use that variable.

// If you don't use fmt. Errorf and simply return a string as an error, it will still work,
// but you lose some benefits that come with using fmt. Errorf.

// The use of error as a return type is a standard convention in Go for functions that may
// produce an error.
