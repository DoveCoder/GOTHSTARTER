package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	gothstarter "gothstarter"
	handlers "gothstarter/handlers"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file:", err)
	}
	router := chi.NewMux()

	router.Handle("/*", gothstarter.Public())
	router.Get("/", handlers.Make(handlers.HandleHome))
	router.Get("/login", handlers.Make(handlers.HandleLoginIndex))

	listenAddr := os.Getenv("LISTEN_ADDR")
	slog.Info("HTTP server listening on", "ListenAddr", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, router))
}
