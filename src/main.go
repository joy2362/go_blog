package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting web server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello your requseted to %v", r.URL.Path)
	})

	log.Fatal(http.ListenAndServe(":8081", nil))
}
