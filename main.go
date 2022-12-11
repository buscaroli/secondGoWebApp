package main

import (
	"fmt"
	"net/http"
)

const port = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Home!")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello About!")
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println("Server up and running on port", port)
	http.ListenAndServe(port, nil)
}
