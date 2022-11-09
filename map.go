package main

import "log"

// Creating my own type of map
type User struct {
	FirstName string
	Lastname string
}


func main() {
	// nameOfMap := make(map[KeyType]ValueType)

	// string map
	// myMap := make(map[string]string)

	// myMap["dog"] = "samson"
	// myMap["other-dog"] = "cassie"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	//int map
	// myMap := make(map[string]int)

	// myMap["First"] = 1
	// myMap["Second"] = 2 
	// log.Println(myMap["First"], myMap["Second"])

	//my own type of map
	myMap := make(map[string]User)

	me :=User {
		FirstName: "Isaac",
		Lastname: "Wanger",
	} 

	myMap["me"] = me
	log.Println(myMap["me"].Lastname)






	
}