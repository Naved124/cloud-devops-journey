package main

import (
	"fmt"
	"net/http"
	"log"
)
func main (){
	// 1. The Router: When a request hits the root URL ("/"), execute this function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Day 2: The server is alive and responding.")
	})


	// 2. The Listener: Bind to port 8080 and keep the process running indefinitely
	fmt.Println("Server listening on port 8080...")
	http.ListenAndServe(":8080", nil)

	// If ListenAndServe throws an error, log.Fatal will print it and crash loudly
	log.Fatal(http.ListenAndServe("8080", nil))

}