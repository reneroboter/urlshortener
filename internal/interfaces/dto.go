package interfaces

import (
	"net/url"
	"regexp"
)

type PostCreateShortURLRequest struct {
	URL string `json:"url"`
}

func (r PostCreateShortURLRequest) Validate() bool {
	parsed, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return false
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return false
	}
	return true
}

type PostCreateShortURLResponse struct {
	Code string `json:"code"`
}

type GetShortURLRequest struct {
	Code string
}

var sha1Regex = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)

func (r GetShortURLRequest) Validate() bool {
	return sha1Regex.MatchString(r.Code)
}

type KafkaEvent struct {
	EventID    string `json:"eventId"`
	Code       string `json:"code"`
	URL        string `json:"url"`
	OccurredAt int64  `json:"occurredAt"`
	Type       string `json:"type"`
}
