package main

import "log"

func main() {

	myVar := "Cars"

	switch myVar {
	case "Cars":
		log.Println("Cars is set to cars")

	case "houses":
		log.Println("houses is set to houses")

	case "phones":
		log.Println("phones is set to phones")

	case "clothes":
		log.Println("clothes is set to clothes")

	default:
		log.Println("I love enjoying the good things of life")
	}
}