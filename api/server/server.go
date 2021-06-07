package server

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"spotifo/api/server/middlewares"
	"spotifo/api/websocket"
)

type Config struct {
	Addr string
}

type APIServer struct {
	*http.Server
	userEndpoint string
	logger       *logrus.Logger
}

func NewAPIServer(cfg Config) *APIServer {
	return &APIServer{
		Server: &http.Server{
			Addr: cfg.Addr,
		},
	}
}

func (s *APIServer) Routers() {
	handler := mux.NewRouter()
	//
	auth := middlewares.Authorization{
		Endpoint: s.userEndpoint,
	}
	//
	handler.Use(auth.Middleware)
	s.Handler = handler
	//
	handler.Handle("/ws", websocket.NewWS(s.logger))
}

func (s *APIServer) Listen() error {
	return s.Server.ListenAndServe()
}

func (s *APIServer) Close() error {
	return s.Server.Close()
}
