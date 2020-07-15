package interfaces

type GetVoiceRequest struct {
	Text string `query:"text"`
}

type GetVoiceResponse struct {
	Content string `json:"content"`
}

