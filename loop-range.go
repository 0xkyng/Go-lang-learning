package main

import "log"

func main() {

	// for i := 0; i <= 10; i++ {
	// 	log.Println(i)
	// }

	/* Lets see what range returns while iterating over different kind of collections in Golang;
     1 .Array or slice: The first value returned in case of array or slice is index and the second value is 
	   element.
     2. String: The first value returned in string is index and the second value is rune int.
     3. Map: The first value returned in map is key and the second value is the value of the key-value pair in 
	   map. */


// using range keyword with for loop to
// iterate over the slice elements

	// animals := []string{"dog", "fish", "horse", "cow", "cat",}
	// for i, animal := range animals {
	// 	log.Println(i, animal)
	// }

// or you can you range over without getting the index by usin the blank identifier(_)
animals := []string{"dog", "fish", "horse", "cow", "cat",}
for _, animal := range animals {
	log.Println(animal)
}

// using range keyword with for loop to
// iterate over maps

// The first value returned in map is key  
// The second value is the value of the key-value pair in map
cars := make(map[string]string)
cars["One"] ="Benz"
cars["Two"] = "BMW"

for carType, cars := range cars {
	log.Println(carType, cars)
}


// using range keyword with for loop to
// iterate over strings

// The first value returned in string is index and 
// The second value is rune int.

firstLine := "I am out running the chariots of the skillful"

for i, line := range firstLine {
	// line is the name of the string
	log.Println(i, line)
}


// using range keyword with for loop to
// iterate over custom type

  type Details struct {
	FirstName string
	LastName  string
	Email     string
	Age       int
  }

  var user []Details
  user = append(user, Details{"Isaac", "Wanger", "codekyng@gmail.com", 30})
  user = append(user, Details{"Pee", "Gee", "peegee@gmail.com", 45})
  user = append(user, Details{"Bright", "Jay", "jay@gmail.com", 25})
  user = append(user, Details{"Gift", "Wee", "gift@gmail.com", 15})

  for _, info := range user{
	log.Println(info.FirstName, info.LastName, info.Email, info.Age)
  }



}