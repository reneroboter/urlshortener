package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/helper"
	"github.com/reneroboter/urlshortener/internal/store"
)

func PostRequestHandler(store store.GeneralStoreInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		hashedUrl := helper.HashUrl(helper.NormalizeUrl(request.Url))

		_, err = store.Get(hashedUrl)
		if err == nil {
			http.Error(w, "url already exists", http.StatusConflict)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		response := PostResponse{
			Code: hashedUrl,
		}

		if err := store.Put(hashedUrl, request.Url); err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func GetRequestHandler(store store.GeneralStoreInterface) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Receive GET request")
		hashedUrl := r.PathValue("hashedUrl")

		if !helper.IsValidSHA1(hashedUrl) {
			http.Error(w, "invalid code", http.StatusBadRequest)
			return
		}

		redirectUrl, err := store.Get(hashedUrl)
		if err != nil {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
	}
}
