package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/helper"
	"github.com/reneroboter/urlshortener/internal/store"
)

func PostRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive POST request")
	decoder := json.NewDecoder(r.Body)
	request := PostRequest{}

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if !helper.IsValidUrl(request.Url) {
		http.Error(w, "invalid URL format", http.StatusBadRequest)
		return
	}

	hashedUrl := helper.HashUrl(request.Url)

	_, ok := store.UrlsMap.Load(hashedUrl)
	if ok {
		http.Error(w, "url already exists", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := PostResponse{
		ID: hashedUrl,
	}
	store.UrlsMap.Store(hashedUrl, request.Url)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetRequestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Receive GET request")
	hashedUrl := r.PathValue("hashedUrl")

	if !helper.IsValidSHA1(hashedUrl) {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	redirectUrl, ok := store.UrlsMap.Load(hashedUrl)
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, redirectUrl.(string), http.StatusMovedPermanently)
}
