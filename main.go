package main

import (
	"fmt"
	"net/http"
	"math/rand"
	"github.com/gorilla/mux"
	"crypto/sha256"
	"encoding/hex"
)

const (
	clientID = "98fc1b94f1e445cebcfe067a505598ba"
	redirectURI = "http://localhost:8080/callback"
	response_type = "code"
	scope = "user-library-read playlist-read-private user-top-read user-read-recently-played"
)

func generateRandomString(length int) []byte{
	text := make([]byte, length)
	var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	for i := 0; i < length; i ++ {
		text[i] += possible[rand.Int() % len(possible)]
	}

	return text
}

func hashSHA256(input []byte) string{
	hash := sha256.New()

	hash.Write([]byte(input))
	hashBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Printing from new thread")
	str := "Hello from go server"
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, str)
}

func login(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "https://google.com", http.StatusFound)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/hello", hello)
	r.HandleFunc("/login", login)
	http.Handle("/", r)
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":5555", nil)
}
