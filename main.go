package main

import (
	"fmt"
	"net/http"
)

// Jika mau deploy harus hapus 127.0.0.1
const portNumber = "127.0.0.1:8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)
}