package main

import "log"

func main() {
	// [] is for strings
	// slice_name := []datatype{values}
	// you can put moe than one thing into it
	mySlice := []string{"Chris", "James", "Steve"}
	log.Println(mySlice)

	// How to add to a slice
	// append adds an element at the ebd of a slice
	mySlice = append(mySlice, "Isaac")

	log.Println(mySlice)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// How to print just the first 3 elements of the slice
	log.Println(numbers[0:3])
}
