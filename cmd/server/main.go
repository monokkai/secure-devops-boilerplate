package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"secure-devops-boilerplate/internal/api"
	"secure-devops-boilerplate/internal/auth"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", api.HelloHandler)

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	mux.HandleFunc("/ready", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "READY")
	})

	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		token, _ := auth.GenerateJWT("admin")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"token": "%s"}`, token)
	})

	secureMux := http.NewServeMux()
	secureMux.HandleFunc("/secure", api.SecureHandler)
	secureMux.Handle("/secure", auth.JWTMiddleware(secureMux))

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Printf("Server starting on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
