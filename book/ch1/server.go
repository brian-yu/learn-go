package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", pong)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func pong(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "PONG!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
