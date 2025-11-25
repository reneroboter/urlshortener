package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
)

var urlsMap = make(map[string]string)

type PostMessage struct {
	Url string
}

type PostResponseMessage struct {
	ID string `json:"id"`
}

func hashUrl(url string) string {
	unhashedUrl := url
	h := sha1.New()
	h.Write([]byte(unhashedUrl))

	hashedUrl := hex.EncodeToString(h.Sum(nil))

	return hashedUrl
}

func isValidSHA1(s string) bool {
	return regexp.MustCompile(`^[a-fA-F0-9]{40}$`).MatchString(s)
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

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		PostMessage := PostMessage{}

		err := decoder.Decode(&PostMessage)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if !isValidUrl(PostMessage.Url) {
			http.Error(w, "invalid URL format", http.StatusBadRequest)
			return
		}

		hashedUrl := hashUrl(PostMessage.Url)

		_, ok := urlsMap[hashedUrl]
		if ok {
			http.Error(w, "url already exists", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := PostResponseMessage{
			ID: hashedUrl,
		}
		urlsMap[hashedUrl] = PostMessage.Url

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("GET /{hashedUrl}", func(w http.ResponseWriter, r *http.Request) {
		hashedUrl := r.PathValue("hashedUrl")

		if !isValidSHA1(hashedUrl) {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		redirectUrl, ok := urlsMap[hashedUrl]
		if !ok {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
	})

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}
