package slackcommands

import (
	"encoding/json"
	"net/http"

	"google.golang.org/appengine/log"
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

	action := payload.Actions[0]

	log.Infof(c, "Request: TeamDomain: %s Action: %s Name: %s Value: %s", payload.Team.Domain, payload.CallbackID, action.Name, action.Value)

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(
		slackCommandResponse{
			ResponseType:   "in_channel",
			DeleteOriginal: true,
			Text:           action.Value,
		},
	)

	return err
}
