package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"google.golang.org/api/option"

	"golang.org/x/oauth2/google"

	"cloud.google.com/go/datastore"
)

var slackClientID = os.Getenv("SLACK_CLIENT_ID")
var slackClientSecret = os.Getenv("SLACK_CLIENT_SECRET")
var datastoreProjectID = "looks-wtf"
var datastoreCreds = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS_JSON")

type slackOauth struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	UserID      string `json:"user_id"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
	Timestamp   time.Time
}

func datastoreClient(ctx context.Context) *datastore.Client {
	jwtConfig, err := google.JWTConfigFromJSON([]byte(datastoreCreds), datastore.ScopeDatastore)
	if err != nil {
		log.Fatalf("Failed to load datastore creds: %v", err)
	}
	ts := jwtConfig.TokenSource(ctx)
	datastoreClient, err := datastore.NewClient(ctx, datastoreProjectID, option.WithTokenSource(ts))
	if err != nil {
		log.Fatalf("Failed to create datastore client: %v", err)
	}
	return datastoreClient
}

func persistSlackOauth(so slackOauth) {
	ctx := context.Background()
	datastoreClient := datastoreClient(ctx)
	so.Timestamp = time.Now().UTC()
	key := datastore.IncompleteKey("slackOauth", nil)
	if _, err := datastoreClient.Put(ctx, key, &so); err != nil {
		log.Fatalf("Failed to persist slackOauth %#v: %v", so, err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")

	if resp, err := http.PostForm(
		"https://slack.com/api/oauth.access",
		url.Values{
			"client_id":     {slackClientID},
			"client_secret": {slackClientSecret},
			"code":          {code},
		},
	); err == nil {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			var msg slackOauth
			err = json.Unmarshal(body, &msg)
			if err == nil && msg.Ok {
				log.Printf("Success, response: %#v", msg)
				fmt.Fprintf(w, "The looks.wtf Slack App has been successfully added!\nGive it a go with the `/look awe` command in Slack.")
				persistSlackOauth(msg)
				return
			}
			log.Printf("Error in response: %#v", string(body))
		}
	}

	fmt.Fprintf(w, "There was an error adding the looks.wth Slack App. Please try again.")
}

func main() {
	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
