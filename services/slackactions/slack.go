package slackcommands

type slackActionPayload struct {
	Actions     []slackActionPayloadAction `json:"actions"`
	Team        slackActionPayloadTeam     `json:"team"`
	CallbackID  string                     `json:"callback_id"`
	ResponseURL string                     `json:"response_url"`
}

type slackActionPayloadTeam struct {
	Domain string `json:"domain"`
}

type slackActionPayloadAction struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

type slackCommandResponse struct {
	ResponseType string `json:"response_type"`
	Text         string `json:"text"`
}
