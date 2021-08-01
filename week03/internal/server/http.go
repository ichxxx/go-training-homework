package server

import (
	"net/http"

	"week03/internal/service"
)

func newServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", service.Echo)
	return mux
}

func NewHttpServer() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: newServerMux(),
	}
}