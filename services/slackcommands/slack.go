package slackcommands

type slackCommandResponse struct {
	ResponseType string                           `json:"response_type"`
	Text         string                           `json:"text"`
	Attachments  []slackCommandResponseAttachment `json:"attachments,omitempty"`
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
