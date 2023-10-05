package main

import (
	// "fmt"
	// "math/rand"
	// "time"
	// "crypto/sha256"
	// "encoding/base64"
	// "bytes"
	"net/http"
	// "net/url"
	// "github.com/gorilla/mux"
)

// This handler will recieve the url from Spotify's API, indicating that the user has been successfully authenticated.
// The url will contain parameters which are parsed, and is how the Autherization Code is recieved.
// The server then sends a message to the frontend, notifying the status of the authentication.
func callbackHandler(w http.ResponseWriter, r *http.Request) {
	// frontEndURL := "http://localhost:3000/home"
	authKey := r.URL.Query().Get("code")	// retrieve the authentication key for the use found in the parameters of the URL

	// req, err := http.NewRequest("POST", frontEndURL, bytes.NewBuffer([]byte("We are authenticated!\nHere is the key: " + authKey)))
	// if err != nil {
	// 	println("unable to create new request :(")
	// }
	// req.Header.Set("Content-Type", "application/json")

	// client := &http.Client{}

	// resp, err := client.Do(req)
	// if err != nil {
	// 	println("unable to send http request :(")
	// }
	// defer resp.Body.Close()

	http.Redirect(w, r, "/home", http.StatusFound)
	println(w, "Authentication successful\nAuth Key: " + authKey)
}


func main() {

	println("Server is now runnning on port 3000!")
	http.HandleFunc("/callback", callbackHandler)
	http.ListenAndServe(":8000", nil)
}