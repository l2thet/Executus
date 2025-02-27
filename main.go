package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	var PORTNUMBER = os.Getenv("ClientPort")
	if PORTNUMBER == "" {
		log.Fatal("ClientPort environment variable not set")
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Create a handler function for the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Printf("Server is running on port %s", PORTNUMBER)
	if err := http.ListenAndServe(PORTNUMBER, nil); err != nil {
		fmt.Println("Server is not running")
	}

}
