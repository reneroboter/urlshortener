package main

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	"github.com/reneroboter/urlshortener/internal/application"
	"github.com/reneroboter/urlshortener/internal/interfaces"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func NewHTTPServer(lc fx.Lifecycle, mux *http.ServeMux) *http.Server {
	srv := &http.Server{Addr: ":8888", Handler: mux}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ln, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return err
			}
			fmt.Println("Starting HTTP server at", srv.Addr)
			go srv.Serve(ln)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})
	return srv
}

func NewServeMux() *http.ServeMux {
	shortURLService := application.NewShortURLService()
	mux := http.NewServeMux()
	mux.HandleFunc("POST /shorten", interfaces.PostCreateShortURLHandler(shortURLService))
	mux.HandleFunc("GET /{code}", interfaces.GetRequestHandler(shortURLService))
	return mux
}

func NewLogger() *slog.Logger {
	return slog.New(slog.NewTextHandler(os.Stdout, nil))
}

func main() {
	slog.Info("Start urlshortener")
	fx.New(
		fx.WithLogger(func(logger *slog.Logger) fxevent.Logger {
			return &fxevent.ConsoleLogger{W: os.Stdout}
		}),
		fx.Provide(NewLogger),
		fx.Provide(NewHTTPServer),
		fx.Provide(NewServeMux),
		fx.Invoke(func(*http.Server) {}),
	).Run()
}
