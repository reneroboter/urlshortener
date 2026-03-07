package application

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashUrl(url string) string {
	unhashedUrl := url
	h := sha1.New()
	h.Write([]byte(unhashedUrl))

	hashedUrl := hex.EncodeToString(h.Sum(nil))

	return hashedUrl
}
