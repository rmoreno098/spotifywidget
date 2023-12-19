package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"github.com/rs/cors"
	"spotify-widget/server/types"
)
var verifier string

func fetchProfile(token string) (string, string, error) {
	url := "https://api.spotify.com/v1/me"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", "", err
	}
	req.Header.Set("Authorization", "Bearer " + token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "", err
	}

	// parse the response body and only return the user's id and display name
	var x types.UserProfile
	err = json.NewDecoder(resp.Body).Decode(&x)
	if err != nil {
		return "", "", err
	}
	resp.Body.Close()

	log.Println("Returning: ", x.ID, x.DisplayName)
	return x.ID, x.DisplayName, nil
}

func verifierHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body) // read the body of the request
	if err != nil {
		log.Println(err)
		return
	}
	defer r.Body.Close()

	var x types.VerResp // create a variable of type verResp
    err = json.Unmarshal(body, &x)	// store body into x
	if err != nil {
		log.Println(err)
		return
	}
	verifier = x.Verifier

	w.WriteHeader(http.StatusOK)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")	// retrieve the code found in the parameters of the callback URL
	if code == "" {
		log.Println("No code found")
		http.Error(w, "No Authorization Code", 1)
		return
	}
	log.Println("Code parameter:", code)

	token := getAccessToken(code, verifier)
	if token == "error" || token == "" {
		log.Println("Fetching access token error")
		http.Error(w, "Error fetching access token", 2)
		return
	}
	log.Println("Token:", token)

	id, name, err := fetchProfile(token)
	if err != nil{
		log.Println("Error fetching profile:", err)
	}

	log.Println("Name:", name)
	log.Println("id:", id)


	// store id and name in the database (token too but needs to be encoded first)

	http.Redirect(w, r, "http://localhost:5173/dashboard", http.StatusFound)
}

func getAccessToken(code string, verifier string) string {
	params := url.Values{
		"client_id":     {"98fc1b94f1e445cebcfe067a505598ba"},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {"http://localhost:8080/callback"},
		"code_verifier": {verifier},
	}
	payload := strings.NewReader(params.Encode())

	resp, err := http.Post("https://accounts.spotify.com/api/token", 
						   "application/x-www-form-urlencoded",
						   payload)
	if err != nil {
		log.Println(err)
		return "error"
	} else {
		defer resp.Body.Close()

		// Read the response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return "error"
		}
		// Print or store the access token (response handling)
		var x types.TokenResp
		err = json.Unmarshal(responseBody, &x)
		if err != nil {
			log.Println(err)
			return "error"
		}

		return x.AccessToken
	}
}

func main() {
	corsHandler := cors.Default()

	log.Println("Server is now runnning on port 8080!")

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(callbackHandler)).ServeHTTP(w, r)
    })

	http.HandleFunc("/verifier", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(verifierHandler)).ServeHTTP(w, r)
    })

	http.ListenAndServe(":8080", nil)
}
