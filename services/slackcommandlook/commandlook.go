package slackcommandlook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/leighmcculloch/looks.wtf/shared/looks"
	"github.com/leighmcculloch/looks.wtf/shared/secrets"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func commandLookHandler(w http.ResponseWriter, r *http.Request) error {
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
	userName := r.FormValue("user_name")
	command := r.FormValue("command")
	tag := r.FormValue("text")
	responseURL := r.FormValue("response_url")

	log.Infof(c, "Request: TeamDomain: %s ChannelName: %s UserName: %s Command: %s Text: %s", teamDomain, channelName, userName, command, tag)

	looksWithTag := looks.LooksWithTag(tag)
	if len(looksWithTag) == 0 {
		fmt.Fprintf(w, "Try using the /look command with one of these words: "+strings.Join(looks.Tags(), ", "))
		return nil
	}

	l := looksWithTag[rand.Intn(len(looksWithTag))]

	client := urlfetch.Client(c)
	body := bytes.Buffer{}
	err := json.NewEncoder(&body).Encode(
		slackCommandResponse{
			ResponseType: "in_channel",
			Text:         l.Plain,
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
