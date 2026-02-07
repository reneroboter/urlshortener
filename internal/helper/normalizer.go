package helper

import (
	"net/url"
	"strings"
)

func NormalizeUrl(raw string) string {
	raw = strings.TrimSpace(raw)

	u, err := url.Parse(raw)
	if err != nil {
		return raw
	}

	u.Scheme = strings.ToLower(u.Scheme)
	u.Host = strings.ToLower(u.Host)

	u.Path = strings.TrimRight(u.Path, "/")

	return u.String()
}
