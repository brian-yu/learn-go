package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	log.Println("Running...")
	log.Fatal(http.ListenAndServe(
		"127.0.0.1:8080",
		http.FileServer(http.Dir(
			filepath.Join(os.Getenv("HOME"), "dev", "go")))))
}
