package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/leighmcculloch/looks.wtf/service/data"
	"github.com/leighmcculloch/looks.wtf/service/shared/secrets"
)

func commandLookHandler(dataLooks map[string][]data.Look, dataTags []string) func(w http.ResponseWriter, r *http.Request) error {
	return func(w http.ResponseWriter, r *http.Request) error {
		c := r.Context()
		defer r.Body.Close()

		slackVerificationToken := secrets.Get(c, "SLACK_VERIFICATION_TOKEN")

		token := r.FormValue("token")
		if token != slackVerificationToken {
			http.Error(w, "401 Unauthorized", http.StatusUnauthorized)
			return nil
		}

		teamDomain := r.FormValue("team_domain")
		channelName := r.FormValue("channel_name")
		userID := r.FormValue("user_id")
		command := r.FormValue("command")
		tag := r.FormValue("text")
		responseURL := r.FormValue("response_url")

		log.Printf("Request: TeamDomain: %s ChannelName: %s UserID: %s Command: %s Text: %s", teamDomain, channelName, userID, command, tag)

		looksWithTag := dataLooks[tag]
		if len(looksWithTag) == 0 {
			fmt.Fprintf(w, "Try using the /look command with one of these words: "+strings.Join(dataTags, ", "))
			return nil
		}

		l := looksWithTag[rand.Intn(len(looksWithTag))]

		client := http.DefaultClient
		body := bytes.Buffer{}
		err := json.NewEncoder(&body).Encode(
			slackCommandResponse{
				ResponseType: "in_channel",
				Text:         fmt.Sprintf("<@%s>: %s", userID, l.Plain),
			},
		)
		if err != nil {
			return fmt.Errorf("Failed to make delayed response post to %s: %s", responseURL, err)
		}
		resp, err := client.Post(responseURL, "application/json", &body)
		if err != nil {
			return fmt.Errorf("Failed to make delayed response post to %s: %s", responseURL, err)
		}
		if resp.StatusCode != 200 {
			return fmt.Errorf("Failed to make delayed response post to %s: status code returned is %d, want 200", responseURL, resp.StatusCode)
		}

		return nil
	}
}
