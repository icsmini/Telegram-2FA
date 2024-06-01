package apiserver

import "net/http"

type APIServer struct {
	config *Config
	Server *http.Server
}

func New(cfg *Config, srv *http.Server) *APIServer {
	return &APIServer{
		config: cfg,
		Server: srv,
	}
}

func (s *APIServer) Start() error {
	return nil
}
