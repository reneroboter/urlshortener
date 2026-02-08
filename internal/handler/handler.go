package handler

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/reneroboter/urlshortener/internal/helper"
	"github.com/reneroboter/urlshortener/internal/store"
	"github.com/twmb/franz-go/pkg/kgo"
)

func PostRequestHandler(store store.GeneralStoreInterface, kafka *kgo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		request := PostRequest{}

		err := decoder.Decode(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		slog.Info("[POST] Store URL", "url", request.Url)

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

		event := KafkaEvent{
			EventID:    uuid.NewString(),
			Code:       hashedUrl,
			URL:        request.Url,
			OccurredAt: time.Now().UnixMilli(),
			Type:       "url_created",
		}

		payload, err := json.Marshal(event)

		if err != nil {
			slog.Error(err.Error())
			return
		}

		record := &kgo.Record{
			Value: payload,
			Topic: "url-shortener-events",
		}

		if err := kafka.ProduceSync(context.Background(), record).FirstErr(); err != nil {
			slog.Error(err.Error())
			return
		}
	}
}

func GetRequestHandler(store store.GeneralStoreInterface, kafka *kgo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.PathValue("code")
		slog.Info("[GET] Request URL", "code", code)

		if !helper.IsValidSHA1(code) {
			http.Error(w, ErrInvalidCode.Error(), http.StatusBadRequest)
			return
		}

		redirectUrl, err := store.Get(code)
		if err != nil {
			http.Error(w, ErrNotFound.Error(), http.StatusNotFound)
			return
		}

		event := KafkaEvent{
			EventID:    uuid.NewString(),
			Code:       code,
			URL:        redirectUrl,
			OccurredAt: time.Now().UnixMilli(),
			Type:       "url_redirected",
		}

		payload, err := json.Marshal(event)

		if err != nil {
			slog.Error(err.Error())
			return
		}

		record := &kgo.Record{
			Value: payload,
			Topic: "url-shortener-events",
		}

		if err := kafka.ProduceSync(context.Background(), record).FirstErr(); err != nil {
			slog.Error(err.Error())
			return
		}

		http.Redirect(w, r, redirectUrl, http.StatusMovedPermanently)
	}
}
