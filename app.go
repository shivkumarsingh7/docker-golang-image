package main

import (
	"fmt"
	"log"
	"net/http"
)

func HomeEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world :)")
}

func main() {
	fmt.Println("server started")
	http.HandleFunc("/", HomeEndpoint)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("server stoped")
}
