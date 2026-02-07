package main

import (
	"fmt"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/handler"
	"github.com/reneroboter/urlshortener/internal/store"
)

func main() {
	store := store.NewTwoLayerStore()
	fmt.Println("two layer")
	fmt.Println("Start urlshortener")
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.PostRequestHandler(store))
	mux.HandleFunc("GET /{hashedUrl}", handler.GetRequestHandler(store))

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}
