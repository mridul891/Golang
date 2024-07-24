package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Strings in Golang")

	data := "apple,orange,banana,Mridul"

	parts := strings.Split(data, ",")
	fmt.Println(parts)

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("count:", count)

	str = "     Hello, Go!"
	trimmed := strings.TrimSpace(str)
	fmt.Println("trimmed :", trimmed)

	//  To explore more about strings package visit go Official Docs

	str1 := "Mridul"
	str2 := "Pandey"
	result := strings.Join([]string{str1, str2 ,"jod"}, " ")
	fmt.Println(result)
}
