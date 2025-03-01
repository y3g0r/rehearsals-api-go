package main

import (
	"fmt"
	"github.com/y3g0r/rehearsals-api-go/internal/config"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Dockerized Go!")
}

func main() {
	cfg := config.LoadConfig()
	// Start the server
	log.Printf("Starting server on %s", cfg.ServerAddress)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(cfg.ServerAddress, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
