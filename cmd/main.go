package main

import (
	"log/slog"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/handler"
	"github.com/reneroboter/urlshortener/internal/store"
)

func main() {
	store := store.NewTwoLayerStore()
	slog.Info("Start urlshortener")
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.PostRequestHandler(store))
	mux.HandleFunc("GET /{code}", handler.GetRequestHandler(store))

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		slog.Error("Something went wrong", "err", err)
		return
	}
}
