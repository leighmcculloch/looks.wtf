package slackcommandlook

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/leighmcculloch/looks.wtf/shared/looks"
	"github.com/leighmcculloch/looks.wtf/shared/secrets"
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

	log.Printf("Request: TeamDomain: %s ChannelName: %s UserName: %s Command: %s Text: %s", teamDomain, channelName, userName, command, tag)

	looksWithTag := looks.LooksWithTag(tag)
	if len(looksWithTag) == 0 {
		fmt.Fprintf(w, "Try using the /look command with one of these words: "+strings.Join(looks.Tags(), ", "))
		return nil
	}

	l := looksWithTag[rand.Intn(len(looksWithTag))]

	w.Header().Add("Content-Type", "application/json")
	response := slackCommandResponse{
		ResponseType: "in_channel",
		Text:         l.Plain,
	}
	enc := json.NewEncoder(w)
	return enc.Encode(response)
}
