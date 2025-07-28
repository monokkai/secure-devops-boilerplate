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
	mux.Handle("/", api.HelloHandler)
}
