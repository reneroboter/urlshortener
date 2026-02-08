package handler

type PostRequest struct {
	Url string `json:"url"`
}

type PostResponse struct {
	Code string `json:"code"`
}

type KafkaEvent struct {
	EventID    string `json:"eventId"`
	Code       string `json:"code"`
	URL        string `json:"url"`
	OccurredAt int64  `json:"occurredAt"`
	Type       string `json:"type"`
}
