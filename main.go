package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/xbmlz/startergo/handlers"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	router := chi.NewMux()

	router.Handle("/*", public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLogin))

	listenAddr := os.Getenv("LISTEN_ADDR")

	slog.Info("HTTP server started", "addr", listenAddr)

	http.ListenAndServe(listenAddr, router)
}
