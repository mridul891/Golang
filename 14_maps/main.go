package main

import "fmt"

func main() {

	fmt.Println("Maps in Golang")

	// Map is a data structure that provides an unordered collection of key-value pairs , where each key must be unique

	//Maps are used to associate values with keys and allow for efficient retrieval of values based on those keys.

	// Some key points about maps:
	// 1. Unordered: Maps in Go are unordered, meaning there is no guaranteed order of
	// key-value pairs when iterating over them.

	// 2. Dynamic Size: The size of a map can grow or shrink dynamically as key-value
	// pairs are added or removed.

	// 3. Keys and Values: Keys and values can be of any comparable type, but all keys
	// must be of the same type, and all values must be of the same type.

	// 4. Initialization: Maps can be initialized using the make function, or using a map literal.

	// creating a map of name with student name and grades
	studentGrades := make(map[string]int)

	studentGrades["Price"] = 34
	studentGrades["ALice"] = 90
	studentGrades["Bob"] = 85
	studentGrades["Charlie"] = 95

	// we cannot insert the data while intializing the map

	// key is very much important it is case sensitive in nature
	fmt.Println("Marks of Bob is :", studentGrades["Bob"])
	studentGrades["Bob"] = 100
	fmt.Println("Marks of Bob is :", studentGrades["Bob"])

	// Delete a key from the map
	delete(studentGrades, "Bob")
	fmt.Println("Marks of Bob is :", studentGrades["Bob"])

	// Checking if key is present or not ,
	// return 2 parameters i.e values=The value which we want , exists=Present in map or not
	grades, exists := studentGrades["David"]
	fmt.Println("Grades of Davide", grades)
	fmt.Println("David exists ", exists)

	// for loop in map
	for index, value := range studentGrades {
		fmt.Printf("Key is %s and marks is %d\n", index, value)
	}

	// Intializing the map with initial value
	person := map[string]int{
		"Alice":   90,
		"Bob":     85,
		"Charlie": 95,
	}
	for index, value := range person {
		fmt.Printf("---Key is %s and marks is %d----\n", index, value)
	}

}
