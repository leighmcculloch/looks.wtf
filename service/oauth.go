package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/leighmcculloch/looks.wtf/service/shared/secrets"
)

type slackOauth struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	UserID      string `json:"user_id"`
	TeamName    string `json:"team_name"`
	TeamID      string `json:"team_id"`
	Timestamp   time.Time
}

func oauthHandler(w http.ResponseWriter, r *http.Request) error {
	c := r.Context()
	code := r.URL.Query().Get("code")
	log.Printf("Request with code: %#v", code)

	slackClientID := secrets.Get(c, "SLACK_CLIENT_ID")
	slackClientSecret := secrets.Get(c, "SLACK_CLIENT_SECRET")

	client := http.DefaultClient
	resp, err := client.PostForm(
		"https://slack.com/api/oauth.access",
		url.Values{
			"client_id":     {slackClientID},
			"client_secret": {slackClientSecret},
			"code":          {code},
		},
	)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var msg slackOauth
	err = json.Unmarshal(body, &msg)
	if err != nil || !msg.Ok {
		log.Printf("Error in response: %#v", string(body))
		return err
	}

	log.Printf("Success, response: %#v", msg)
	fmt.Fprintf(w, "The looks.wtf Slack App has been successfully added!\nGive it a go with the `/look awe` command in Slack.")
	return nil
}
