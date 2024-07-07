package spotify

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.

var ch = make(chan *spotify.Client)

func storeToken(addr, clientId, clientSecret, state string) {
	redirectUrl := fmt.Sprintf("http://%v/callback", addr)
	if strings.HasPrefix(addr, ":") {
		redirectUrl = fmt.Sprintf("http://localhost:%v/callback", addr)
	}
	auth := spotifyauth.New(
		spotifyauth.WithClientID(clientId),
		spotifyauth.WithClientSecret(clientSecret),
		spotifyauth.WithRedirectURL(redirectUrl),
		spotifyauth.WithScopes(spotifyauth.ScopeUserReadPlaybackState, spotifyauth.ScopeUserModifyPlaybackState))
	// first start an HTTP server
	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		completeAuth(w, r, auth, state)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	user, err := client.CurrentUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("You are logged in as:", user.ID)
}

func completeAuth(w http.ResponseWriter, r *http.Request, auth *spotifyauth.Authenticator, state string) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Fatal(err)
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Fatalf("State mismatch: %s != %s\n", st, state)
	}

	// use the token to get an authenticated client
	client := spotify.New(auth.Client(r.Context(), tok))
	fmt.Fprintf(w, "Login Completed!")
	f, err := os.Create("token.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if err := json.NewEncoder(f).Encode(tok); err != nil {
		log.Fatal(err)
	}
	ch <- client
}
