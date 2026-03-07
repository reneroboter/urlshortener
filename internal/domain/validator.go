package domain

import (
	"net/url"
	"regexp"
)

var sha1Regex = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)

func IsValidCode(Code string) bool {
	return sha1Regex.MatchString(Code)
}

func IsValidURL(URL string) bool {
	parsed, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return false
	}
	return true
}
