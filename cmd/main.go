package main

import (
	"log/slog"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/application"
	"github.com/reneroboter/urlshortener/internal/infrastructure"
	"github.com/reneroboter/urlshortener/internal/interfaces"
)

var Store = infrastructure.NewShortUrlRepository()

func main() {

	shortURLService := application.NewShortURLService()
	slog.Info("Start urlshortener")
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", interfaces.PostCreateShortURLHandler(shortURLService))
	mux.HandleFunc("GET /{code}", interfaces.GetRequestHandler(shortURLService))

	err = http.ListenAndServe(":8888", mux)
	if err != nil {
		slog.Error("Something went wrong", "err", err)
		return
	}
}
