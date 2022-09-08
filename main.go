package main

import (
	"errors"
	"fmt"
	"net/http"
)

// Jika mau deploy harus hapus 127.0.0.1
const portNumber = "127.0.0.1:8080"

// Ini adalah handlerfunction
// Ingat bahwa handlerfunction HARUS handle two parameters, responsewriter dan request
func Home(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request){
	sum := addValues(2,2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("2+2 is %d", sum))
}

// what if I pass x and y that is not integer?
// kita handle error dengan tambah return error
func addValues (x,y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request){
	// kita ga ingin melakukan pembagian dengan 0
	f, err := divideValues(100.0, 0.0)
	if err != nil {
		fmt.Fprintf(w, "cannot divide by zero")
		return
	}
	fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", 100.0, 0.0,f))

}

func divideValues(x,y float32) (float32,error){
	if y <= 0 {
		err := errors.New("cannot divide by zero")
		return 0, err
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}