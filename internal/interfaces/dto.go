package interfaces

import (
	"github.com/reneroboter/urlshortener/internal/domain"
)

type PostCreateShortURLRequest struct {
	URL string `json:"url"`
}

func (r PostCreateShortURLRequest) Validate() bool {
	return domain.IsValidURL(r.URL)
}

type PostCreateShortURLResponse struct {
	Code string `json:"code"`
}

type GetShortURLRequest struct {
	Code string
}

func (r GetShortURLRequest) Validate() bool {
	return domain.IsValidCode(r.Code)
}

type KafkaEvent struct {
	EventID    string `json:"eventId"`
	Code       string `json:"code"`
	URL        string `json:"url"`
	OccurredAt int64  `json:"occurredAt"`
	Type       string `json:"type"`
}
