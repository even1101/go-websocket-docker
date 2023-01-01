package main

import (
	"net/http"

	"github.com/bmizerany/pat"

	"example.com/demo/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux

}
