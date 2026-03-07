package domain

import (
	"net/url"
	"regexp"
)

type ShortURL struct {
	Code string `json:"code"`
	URL  string `json:"url"`
}

var sha1Regex = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)

func (s ShortURL) isValidCode() bool {
	return sha1Regex.MatchString(s.Code)
}

func (s ShortURL) isValidURL() bool {
	parsed, err := url.ParseRequestURI(s.URL)
	if err != nil {
		return false
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return false
	}
	return true
}

func (s ShortURL) validate() {
	s.isValidURL()
	s.isValidCode()
}
