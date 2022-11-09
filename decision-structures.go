package main

import "log"

func main() {
	// isTrue := "True"
	//or
	var isTrue bool

	isTrue = true

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is ", false)

	}

	myNum := 100
	isYes := false

	if myNum > 99 && !isYes {
		log.Println("myNum is greater than 99 and isYes is set to true")
	}

}