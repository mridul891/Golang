package main

import "fmt"

// Define a struct named Person
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}

//  Yeah sab tarike ke struct ek type of data type ban jaate hia jisse hum kahi bhi use kar sakte hai

func main() {

	fmt.Println("Struct in GoLang")

	// In go  a struct is a composite data type that groups together variables ( fields or members ) under a single name.
	// Each field in a struct can have a different data types , and structs are used to create more complex data structure.
	// They play a crucial role in creating complex types and managing relationships betwen different pieces of data.

	// simple intialization
	var prince Person
	fmt.Println("Prince Person :", prince)
	prince.FirstName = "Prince"
	prince.LastName = "Pandey"
	prince.Age = 24

	fmt.Println("Prince Person :", prince)

	// 2nd way

	person1 := Person{
		FirstName: "mridul",
		LastName:  "Pandey",
		Age:       20,
	}
	fmt.Println("person1 Person :", person1)

	// using new keyword
	var person2 = new(Person)
	person2.FirstName = "Bhumika"
	person2.LastName = "Pandey"
	person2.Age = 24
	fmt.Println("person2 Person :", person2)
	// here & came because person2 act as an pointer

	var employee1 Employee

	employee1.Person_Details = Person{
		FirstName: "Mridul Shansha",
		LastName:  "Pandey",
		Age:       21,
	}
	employee1.Person_Address = Address{
		House: 12,
		Area:  "Vaishale",
		State: "UP",
	}
	employee1.Person_Contact = Contact{
		Email: "mrid@gmail.com",
		Phone: "9829323292",
	}

	fmt.Println("Employee 1:", employee1)
}
