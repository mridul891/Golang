// In go , packages are used instead of classes . There is no concept like OOPs in java lang;
package main

import (
	"fmt"
	"mylearning/myutil"
)

// Go lang mai ek folder  mai jitni bhi ile hai wo ab same package ko import karengi koi hume dusri package ko import karna hai toh hume uske liya new folder bana padhega
func main() {
	fmt.Println("Started the import ")

	myutil.Printmessage("hello mridul from Public function")

}
