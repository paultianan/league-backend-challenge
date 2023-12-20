package main

import (
	"fmt"
	"net/http"
	"paul-tianan/league-backend-challenge/handlers"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/file-path/matrix.csv' "localhost:8080/echo"

const PORT = 8080

func main() {
	http.HandleFunc("/echo", handlers.EchoHandler)
	http.HandleFunc("/invert", handlers.InvertHandler)
	http.HandleFunc("/flatten", handlers.FlattenHandler)
	http.HandleFunc("/sum", handlers.SumHandler)
	http.HandleFunc("/multiply", handlers.MultiplyHandler)

	fmt.Printf("Server is running on port %d...\n", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)
	if err != nil {
		fmt.Println("Error starting the server:", err)
	}
}
