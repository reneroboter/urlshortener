package helper

import (
	"crypto/sha1"
	"encoding/hex"
	"net/url"
	"regexp"
)

var sha1Regex = regexp.MustCompile(`^[a-fA-F0-9]{40}$`)

func IsValidSHA1(s string) bool {
	return sha1Regex.MatchString(s)
}

func IsValidUrl(u string) bool {
	parsed, err := url.ParseRequestURI(u)
	if err != nil {
		return false
	}
	if parsed.Scheme == "" || parsed.Host == "" {
		return false
	}
	return true
}

func HashUrl(url string) string {
	unhashedUrl := url
	h := sha1.New()
	h.Write([]byte(unhashedUrl))

	hashedUrl := hex.EncodeToString(h.Sum(nil))

	return hashedUrl
}
