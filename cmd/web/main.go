package main

import (
	"fmt"
	"net/http"

	"github.com/buscaroli/secondGoWebApp/pkg/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println("Server up and running on port", port)
	http.ListenAndServe(port, nil)
}
