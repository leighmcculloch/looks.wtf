package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func actionHandler(w http.ResponseWriter, r *http.Request) error {
	defer r.Body.Close()

	payloadJSON := r.FormValue("payload")
	payload := slackActionPayload{}
	err := json.Unmarshal([]byte(payloadJSON), &payload)
	if err != nil {
		return err
	}

	action := payload.Actions[0]

	log.Printf("Request: TeamDomain: %s Action: %s Name: %s Value: %s", payload.Team.Domain, payload.CallbackID, action.Name, action.Value)

	w.Header().Add("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(
		slackCommandResponse{
			ResponseType:   "in_channel",
			DeleteOriginal: true,
			Text:           fmt.Sprintf("<@%s>: %s", payload.User.ID, action.Value),
		},
	)

	return err
}
