package interfaces

import (
	"encoding/json"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/application"
)

func PostCreateShortURLHandler(s application.ShortURLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		request := PostCreateShortURLRequest{}

		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if !request.Validate() {
			http.Error(w, "invalid URL format", http.StatusBadRequest)
			return
		}

		shortURL := s.CreateShortURL(request.URL)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		response := PostCreateShortURLResponse{
			Code: shortURL.Code,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func GetRequestHandler(s application.ShortURLService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		request := GetShortURLRequest{
			Code: r.PathValue("code"),
		}

		if !request.Validate() {
			http.Error(w, ErrInvalidCode.Error(), http.StatusBadRequest)
			return
		}

		err, shortURL := s.ReturnShortURL(request.Code)

		if err != nil {
			http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, shortURL.URL, http.StatusMovedPermanently)
	}
}
