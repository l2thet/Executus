package main

import (
	"fmt"
	"net/http"
)

const PORTNUMBER = 8081

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Create a handler function for the root URL
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	fmt.Printf("Server is running on port %d", PORTNUMBER)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORTNUMBER), nil); err != nil {
		fmt.Println("Server is not running")
	}

}
