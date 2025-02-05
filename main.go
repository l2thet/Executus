package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/listavailablemusic", listAvailableMusic)

	mux.HandleFunc("/api/servesong", serveSong)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Printf("Starting server on %s", server.Addr)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func listAvailableMusic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusOK)

		files, err := readDir("./assets/music")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}

}

func readDir(dirname string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, info.Name())
		}
		return nil
	})
	return files, err
}

func serveSong(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		fileName := r.URL.Query().Get("name")
		if fileName == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		http.ServeFile(w, r, "./assets/music/"+fileName)
	}
}
