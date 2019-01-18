package service

type slackCommandResponse struct {
	ResponseType   string                           `json:"response_type"`
	DeleteOriginal bool                             `json:"delete_original,omitempty"`
	Text           string                           `json:"text"`
	Attachments    []slackCommandResponseAttachment `json:"attachments,omitempty"`
}

type slackCommandResponseAttachment struct {
	Text       string                                 `json:"text"`
	Fallback   string                                 `json:"fallback"`
	CallbackID string                                 `json:"callback_id"`
	Actions    []slackCommandResponseAttachmentAction `json:"actions,omitempty"`
}

type slackCommandResponseAttachmentAction struct {
	Name  string `json:"name"`
	Text  string `json:"text"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

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
