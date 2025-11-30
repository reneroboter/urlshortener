package main

import (
	"net/url"
	"regexp"
)

var sha1Regex = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)

func isValidSHA1(s string) bool {
	return sha1Regex.MatchString(s)
}

func isValidUrl(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return false
	}
	return true
}
