package main

import (
	"fmt"
	"net/http"
)

func main() {
	
	// http.HandleFunc to tell the server which function to call
	// to handle a request to the server

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// Fprintf formats according to a format specifier and writes to w. 
		// And it returns the number of bytes written and any write error encountered
		// n int for(number of bytes written)
		// err error for(error encoutered)

		n, err := fmt.Fprintf(w, "Hello, world!")
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Sprintf() allows you to take different data types
		// and return them as a string

		fmt.Println(fmt.Sprintf("Number of bytes written: %d", n))
		
	}) 

	// http.ListenAndServe function to start the server 
	// and tell it to listen for new HTTP requests and 
	// then serve them using the handler functions you set up.
	_= http.ListenAndServe(":8080", nil)
}