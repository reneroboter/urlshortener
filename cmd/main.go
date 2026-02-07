package main

import (
	"fmt"
	"net/http"

	"github.com/reneroboter/urlshortener/internal/handler"
)

func main() {
	fmt.Println("Start urlshortener")
	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", handler.PostRequestHandler)
	mux.HandleFunc("GET /{hashedUrl}", handler.GetRequestHandler)

	err := http.ListenAndServe(":8888", mux)
	if err != nil {
		fmt.Println(err)
		return
	}
}
