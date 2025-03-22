package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", greet)

	fmt.Println("Starting server on :4000")

	http.ListenAndServe(":4000", mux)
}
