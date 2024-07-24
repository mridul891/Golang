package main

import (
	"fmt"
	"time"
)

func main() {

	// Go Provides a powerful time package for handling time and date related operations.
	// The format "2006-01-02 15:04:05"
	// In Go time package the refernce date and time for formatting are fixes as "january 2 ,2006 . 3:04:05 PM ,MST (Mountain standard time)"
	// This time is choosen because it is easy to remember and it is also the time when Go was officially announced
	// 2006-  year 01- month 02 - day 15-hrs 04 -minutes 05 - seconds

	currentTime := time.Now()
	fmt.Println("Current time :", currentTime)
	fmt.Printf("Type of the current Time is %T\n", currentTime)

	formatted := currentTime.Format("02-01-2006 ,15:04:05 Monday")
	fmt.Println("Formatted time :", formatted)

	// while aprsing the datastr we always have to create a  layout_str  to parse in the function and make it successfully formated
	layout_str := "2006-01-02"
	dateStr := "2024-11-23"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted time :", formatted_time)

	// add 1 more daty in currentTime
	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println(" new_date:", new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("formatted_new_date :", formatted_new_date)

}
