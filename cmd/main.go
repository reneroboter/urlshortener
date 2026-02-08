package main

import (
	"log/slog"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/handler"
	"github.com/reneroboter/urlshortener/internal/store"
	kafka "github.com/reneroboter/urlshortener/pkg/kafka"
)

func main() {
	store := store.NewTwoLayerStore()
	kafkaClient, err := kafka.NewKafkaClient()
	if err != nil {
		slog.Error("Something went with kafka", "err", err)
		panic(err)
		return
	}
	slog.Info("Start urlshortener")
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.PostRequestHandler(store, kafkaClient))
	mux.HandleFunc("GET /{code}", handler.GetRequestHandler(store, kafkaClient))

	err = http.ListenAndServe(":8888", mux)
	if err != nil {
		slog.Error("Something went wrong", "err", err)
		return
	}
}
