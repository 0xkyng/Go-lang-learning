
// Go program to illustrate the
// method with struct type receiver
package main

import "fmt"

// Author structure
type author struct {
	name	 string
	branch string
	particles int
	salary int
}

// Method with a receiver
// of author type
func (a author) show() {

	fmt.Println("Author's Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Published articles: ", a.particles)
	fmt.Println("Salary: ", a.salary)
}

// Main function
// func main() {

// 	// Initializing the values
// 	// of the author structure
// 	res := author{
// 		name:	 "Sona",
// 		branch: "CSE",
// 		particles: 203,
// 		salary: 34000,
// 	}

// 	// Calling the method
// 	res.show()
// }


// 


type Employee struct {  
    name string
    age  int
}

/*
Method with value receiver  
*/
func (e Employee) changeName(newName string) {  
    e.name = newName
}

/*
Method with pointer receiver  
*/
func (e *Employee) changeAge(newAge int) {  
    e.age = newAge
}

func main() {  
    e := Employee{
        name: "Mark Andrew",
        age:  50,
    }
    fmt.Printf("Employee name before change: %s", e.name)
    e.changeName("Michael Andrew")
    fmt.Printf("\nEmployee name after change: %s", e.name)

    fmt.Printf("\n\nEmployee age before change: %d", e.age)
    (&e).changeAge(51)
    fmt.Printf("\nEmployee age after change: %d", e.age)
}







// func main() {
// 	d, err := divide(5.0, 2.0)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(d)

// }


// func divide(a, b float64) (float64, error) {
// 	if b == 0.0 {
// 		return 0.0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }
