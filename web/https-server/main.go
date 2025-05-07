package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello This is Gane Solutions")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Starting HTTPS server at https://localhost:8443")
	err := http.ListenAndServeTLS(":8443", "cert.pem", "key.pem", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
