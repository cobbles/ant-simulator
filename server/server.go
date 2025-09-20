package server

import (
	"github.com/cobbles/ant-simulator/app"
	"net/http"
)

func New() *http.Server {
	app.Start()
	mux := http.NewServeMux()
	registerRoutes(&app.World, mux)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}
