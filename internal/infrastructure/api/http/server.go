package httpserver

import (
	"log"
	"net/http"
)

type Server struct {
	Port string
}

func (s *Server) Start() {

	server := &http.Server{
		Addr:    ":" + s.Port,
		Handler: nil,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed %v", err)
	}
}
