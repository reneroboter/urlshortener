package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

var urlsMap = sync.Map{}

type PostRequest struct {
	Url string `json:"url"`
}

type PostResponse struct {
	ID string `json:"id"`
}

func PostRequestHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	request := PostRequest{}

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !isValidUrl(request.Url) {
		http.Error(w, "invalid URL format", http.StatusBadRequest)
		return
	}

	hashedUrl := hashUrl(request.Url)

	_, ok := urlsMap.Load(hashedUrl)
	if ok {
		http.Error(w, "url already exists", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := PostResponse{
		ID: hashedUrl,
	}
	urlsMap.Store(hashedUrl, request.Url)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetRequestHandler(w http.ResponseWriter, r *http.Request) {
	hashedUrl := r.PathValue("hashedUrl")

	if !isValidSHA1(hashedUrl) {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	redirectUrl, ok := urlsMap.Load(hashedUrl)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, redirectUrl.(string), http.StatusMovedPermanently)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", PostRequestHandler)
	mux.HandleFunc("GET /{hashedUrl}", GetRequestHandler)

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}
