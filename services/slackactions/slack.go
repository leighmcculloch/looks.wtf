package slackcommands

type slackActionPayload struct {
	Actions     []slackActionPayloadAction `json:"actions"`
	Team        slackActionPayloadTeam     `json:"team"`
	User        slackActionPayloadUser     `json:"user"`
	CallbackID  string                     `json:"callback_id"`
	ResponseURL string                     `json:"response_url"`
}

type slackActionPayloadTeam struct {
	Domain string `json:"domain"`
}

type slackActionPayloadUser struct {
	ID string `json:"id"`
}

type slackActionPayloadAction struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type slackCommandResponse struct {
	ResponseType   string `json:"response_type"`
	DeleteOriginal bool   `json:"delete_original"`
	Text           string `json:"text"`
}
