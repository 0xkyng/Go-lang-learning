package main

import "fmt"

func pointers() {
	//POINTERS
	var myString string
	myString = "Green"

	// fmt.Println("myString is set to", myString)

	fmt.Println("myString is set to", myString)
	changeUsingPointer(&myString) // whwn you wanto reference a variable, you put & in front of it
	fmt.Println(" after func calls myString is set to", myString)

}

func changeUsingPointer(s *string) {
	fmt.Println("s is set to", s)
	newValue := "Red"  // This is a shortcout for creating and storing variables in one line of code
	*s = newValue // Red which is the value of red
}