package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, what's your name?")
	// var name string

	// // scan method only reads until the first white space character
	// fmt.Scan(&name)
	// fmt.Println("Hello , Mr.", name)

	// bufio.NewReader(os.Stdin) creates a new buffered reader that reads from the standard input(os.Stdin)
	// os = operating system , Stdin = standard input
	reader := bufio.NewReader(os.Stdin)
	// reads the input from the user until a  new line is  provided by the user
	Fullname, _ := reader.ReadString('\n')
	fmt.Println("Hello , Mr.", Fullname)

	// A buffered reader isa type od readers taht reads data from an underlying source , such as a file or standard input from the keyboard . and stores that data in a buffer . The buffer is a temporary storage area in memory . Buffered readers are commonly used to improve the efficieny of input operators.
}
