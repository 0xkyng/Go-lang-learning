package main

import "fmt"

// An interface is a custom type that is used to specify
// a set of one or more method signatures and the interface is abstract,
// so you are not allowed to create an instance of the interface.

// Syntax
/*
   type interface_name interface{
      // Method signatures
     }
*/

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}


func main() {
	dog := Dog{
		Name: "Upkeep",
		Breed: "German shepherd",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name: "Junk",
		Color: "grey",
		NumberOfTeeth: 38,
	}

	PrintInfo(&gorilla)

}


func PrintInfo(name Animal) {
	fmt.Println("This animal says", name.Says(), "and has", name.NumberOfLegs(), "legs")
}


func (say *Dog) Says() string {
	return "wolf"

}

func (legs *Dog) NumberOfLegs() int {
	return 4
}


func (say *Gorilla) Says() string {
	return "Ahh"

}

func (legs *Gorilla) NumberOfLegs() int {
	return 2
}