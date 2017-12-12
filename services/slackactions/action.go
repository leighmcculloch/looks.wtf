package slackcommands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func actionHandler(w http.ResponseWriter, r *http.Request) error {
	c := r.Context()
	defer r.Body.Close()

	payloadJSON := r.FormValue("payload")
	payload := slackActionPayload{}
	err := json.Unmarshal([]byte(payloadJSON), &payload)
	if err != nil {
		return err
	}

	responseURL := payload.ResponseURL
	action := payload.Actions[0]

	log.Infof(c, "Request: TeamDomain: %s Action: %s Name: %s Value: %s", payload.Team.Domain, payload.CallbackID, action.Name, action.Value)

	client := urlfetch.Client(c)
	body := bytes.Buffer{}
	err = json.NewEncoder(&body).Encode(
		slackCommandResponse{
			ResponseType: "in_channel",
			Text:         action.Value,
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
