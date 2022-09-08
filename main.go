package main

import (
	"fmt"
	"net/http"
)

func main() {
	// argument untuk HandleFunc adalah: path name yang akan di listen to, sebuah function yang akan dijalankan
	// function ini menerima method bawaan ResponseWriter, dan juga pointer (* => menandakan pointer, yaitu address in memory where some value is stored)
	// kenapa harus pake pointer? Karena the HandleFunc requires the request to be a pointer to a request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		n, err := fmt.Fprintf(w, "Hello World")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})

	// kita buat dia serve request
	// argumentnya adalah port dan handler. Handler kita gausah masukkan krn akan di handle HandleFunc
	// arti dari _ = adalah ignore error
	_ = http.ListenAndServe(":8080", nil)
}