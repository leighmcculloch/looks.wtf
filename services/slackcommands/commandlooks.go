package slackcommands

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"

	"github.com/leighmcculloch/looks.wtf/shared/looks"
	"github.com/leighmcculloch/looks.wtf/shared/secrets"
	"google.golang.org/appengine/log"
)

func commandLooksHandler(w http.ResponseWriter, r *http.Request) error {
	c := r.Context()
	defer r.Body.Close()

	slackVerificationToken := secrets.Get(c, "SLACK_VERIFICATION_TOKEN")
	log.Infof(c, slackVerificationToken)

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

	log.Infof(c, "Request: TeamDomain: %s ChannelName: %s UserName: %s Command: %s Text: %s", teamDomain, channelName, userName, command, tag)

	looksWithTag := looks.LooksWithTag(tag)
	if len(looksWithTag) == 0 {
		fmt.Fprintf(w, "Try using the /look command with one of these words: "+strings.Join(looks.Tags(), ", "))
		return nil
	}

	maxLooks := 5
	if maxLooks > len(looksWithTag) {
		maxLooks = len(looksWithTag)
	}

	actions := []slackCommandResponseAttachmentAction{}
	for i := 0; i < maxLooks; i++ {
		lookIdx := rand.Intn(len(looksWithTag))

		l := looksWithTag[lookIdx]
		actions = append(
			actions,
			slackCommandResponseAttachmentAction{
				Name:  "look",
				Text:  l.Plain,
				Type:  "button",
				Value: l.Plain,
			},
		)

		looksWithTag = append(looksWithTag[:lookIdx], looksWithTag[lookIdx+1:]...)
	}

	w.Header().Add("Content-Type", "application/json")
	response := slackCommandResponse{
		ResponseType: "ephemeral",
		Text:         fmt.Sprintf("There are several looks tagged with `%s`.", tag),
		Attachments: []slackCommandResponseAttachment{
			{
				Text:       "Choose a look",
				Fallback:   "Oh no, something went wrong",
				CallbackID: "looks",
				Actions:    actions,
			},
		},
	}
	enc := json.NewEncoder(w)
	return enc.Encode(response)
}
