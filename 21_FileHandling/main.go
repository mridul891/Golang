package main

import (
	"fmt"
	"io"
	"os"
)

func ReadingwithBuffer() {

	fmt.Println("Reading a File in Golang using Buffer ")

	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while Opening file", err)
		return
	}
	defer file.Close()

	// Create a Buffer to read the file content

	// make([]byte , 1024 ) creates a byte slice (buffer) with a capacity of 1024 byte . This buffer will be used to read chunks of the file

	buffer := make([]byte, 1024)

	// Read the file

	for {
		n, err := file.Read(buffer)
		// file.read(buffer) reads content from the file into the buffer. The loop continues till the EOF is reached
		// EOF = End of File
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error while reading the file")
			return
		}

		// fmt.Println(string(buffer[:n])) processes and prints the read content , In  this example it prints the content as a string

		fmt.Println(string(buffer[:n]))
	}
}

func Write() {
	// Writting in a file

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error Creating file", err)
		return
	}
	defer file.Close()

	content := "Hello by Mridul Pandey the jod content creator"
	byte, error := io.WriteString(file, content+"\n")
	fmt.Println("Byte written", byte)
	if error != nil {
		fmt.Println("Error while writing file :", error)
	}

	// io.WriteString returns us 2 variables :

	// 1: No of bytes written :This is an integer representing the number of bytes written to the file .
	// 2: Error Present while creating a file : if the write opreation encounters an error , it will be non-nil indicating that something went wrong
}

func Readingwithioutils() {
	fmt.Println("Reading file using io.utils function ")

	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	fmt.Println(string(content))

	// ioutil.ReadFile("example.txt") reads the entire content of the file into a byte slice ( content)

	// This method is convenient for scenerios where you want to read the entire content  of a file into memory .
	// Not suitable for latge file as it loads the entire memory at onve

}
func main() {

	fmt.Println("File Handling in Golang")

	// File operations inGo involve working with the os and io/ioutil packages . Here's a basic guide to common file operations like creating, reading and writing file

	// file.Close() is important in manu cases when you are done working with a file . When you open a file using function like os.Create , os.Open or others. You are acquiring system resources to interact with that file . Failing to close the file properly can lead to resource leaks and might cause issues like running out of file descriptors , The close() Function is responsible for releasing those resources .

	Write()

	fmt.Println("Successfully created file")

	ReadingwithBuffer()
	Readingwithioutils()
}
