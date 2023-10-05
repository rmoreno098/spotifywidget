package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Printing from new thread")
	str := "Hello from go server"
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, str)
	time.Sleep(5 * time.Second)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/hello", hello)

	http.Handle("/", r)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
