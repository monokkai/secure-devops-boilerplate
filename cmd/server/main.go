package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"secure-devops-boilerplate/internal/api"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HelloHandler)

	fmt.Printf("Server running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
