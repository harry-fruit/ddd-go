package httpserver

import (
	"log"
	"net/http"
)

type HTTPServer struct {
	Port string
}

func (s *HTTPServer) Start() {

	server := &http.Server{
		Addr:    ":" + s.Port,
		Handler: nil,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server failed %v", err)
	}
}

func NewHTTPServer(port string) *HTTPServer {
	return &HTTPServer{
		Port: port,
	}
}
