package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"log"
	"net/url"
	"strings"
	"github.com/rs/cors"
	"spotify-widget/server/types"
	"github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/apis"
    "github.com/pocketbase/pocketbase/core"
	// "os"
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

	return x.ID, x.DisplayName, nil
}

func verifierHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)	// read the body of the request
	if err != nil {
		fmt.Println(err)
		return
	}
	defer r.Body.Close()

	var x types.VerResp	// create a variable of type verResp
    err = json.Unmarshal(body, &x)	// store body into x
	if err != nil {
		fmt.Println(err)
		return
	}
	verifier = x.Verifier	// store the verifier into the global variable

	w.WriteHeader(http.StatusOK)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")	// retrieve the code found in the parameters of the callback URL
	if code == "" {
		fmt.Println("Authentication failed")
		return
	}

	token := getAccessToken(code, verifier)
	if token == "error" {
		fmt.Println("Authentication failed")
		return
	}

	id, name, err := fetchProfile(token)
	if err != nil {
		log.Println(err)
	}
		
	res, err := app.Dao().DB().
		NewQuery("CREATE INDEX name_idx ON users (name)").
		Execute()

	// store this in the database
	

	http.Redirect(w, r, "http://localhost:5173/dashboard", http.StatusFound)
}

// func profileHandler(w http.ResponseWriter, r *http.Request) {
// 	token := fetchToken()
// 	resp, err := fetchProfile(token)
// 	if err != nil {
// 		fmt.Println("Error fetching profile")
// 		fmt.Println(err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	rawJSON, err := io.ReadAll(resp.Body)
//     if err != nil {
//         http.Error(w, "Error reading Spotify response", http.StatusInternalServerError)
//         return
//     }

//     // Send the raw JSON data to the frontend
//     w.Header().Set("Content-Type", "application/json")
//     w.Write(rawJSON)
// }

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
		fmt.Println(err)
		return "error"
	} else {
		defer resp.Body.Close()

		// Read the response body
		responseBody, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
			return "error"
		}
		// Print or store the access token (response handling)
		var x types.TokenResp
		err = json.Unmarshal(responseBody, &x)
		if err != nil {
			fmt.Println(err)
			return "error"
		}

		return x.AccessToken
	}
}

func main() {
	app := pocketbase.New()
	// app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
    //     e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))
    //     return nil
    // })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }

	corsHandler := cors.Default()

	fmt.Println("Server is now runnning on port 8080!")

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(callbackHandler)).ServeHTTP(w, r)
    })

	http.HandleFunc("/verifier", func(w http.ResponseWriter, r *http.Request) {
        corsHandler.Handler(http.HandlerFunc(verifierHandler)).ServeHTTP(w, r)
    })

	// http.HandleFunc("/getProfile", func(w http.ResponseWriter, r *http.Request) {
	// 	corsHandler.Handler(http.HandlerFunc(profileHandler)).ServeHTTP(w, r)
	// })

	http.ListenAndServe(":8080", nil)
}