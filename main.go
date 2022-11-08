package main

import "fmt"

func main() {
	// fmt.Println("Hello, world!")

	// var whatToSay string 

	// var integer int

	// whatToSay = "Goodbye, good world"

	// fmt.Println(whatToSay)

	// integer = 7
	// fmt.Println("integer is set to", integer)

	// whatWasSaid, theOtherThingThatWasSaid := saySomething()
	// fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)


	//POINTERS
	var myString string
	myString = "Green"

	// fmt.Println("myString is set to", myString)

	fmt.Println("myString is set to", myString)
	changeUsingPointer(&myString) // whwn you wanto reference a variable, you put & in front of it
	fmt.Println(" after func calls myString is set to", myString)
	
};

// POINTERS
// when you want to create  or reference an actual pointer, you put * in front of it
func changeUsingPointer(s *string) {
	fmt.Println("s is set to", s)
	newValue := "Red"  // This is a shortcout for creating and storing variables in one line of code
	*s = newValue // Red which is the value of red
}


// func saySomething() (string, string) {
// 	return "something", "else"
// }