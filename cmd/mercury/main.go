package main

import (
	"fmt"
	"log"
	"net/http"
)


func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Mercury email service is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Mercury Email Service")
}