package main

import "fmt"

type Address struct {
	city   string
	state  string
	street string
}

type Person struct {
	name    string
	age     int
	sex     string
	address Address
}

func main() {
	p1 := Person{
		name: "Isaac",
		age:  26,
		sex:  "male",
		address: Address{
			city:   "Lagos",
			state:  "Lagos",
			street: "Taike street",
		},
	}
	fmt.Println("Name", p1.name)
    fmt.Println("Sex", p1.sex)
    fmt.Println("Age", p1.age)
    fmt.Println("City", p1.address.city)
    fmt.Println("State", p1.address.state)
    fmt.Println("Street", p1.address.street)

}