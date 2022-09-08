package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Ini adalah handlerfunction
// Ingat bahwa handlerfunction HARUS handle two parameters, responsewriter dan request
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	sum := AddValues(2,2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("2+2 is %d", sum))
}

// what if I pass x and y that is not integer?
// kita handle error dengan tambah return error
func AddValues (x,y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}