package main

import (
	"fmt"
	"log"

	"github.com/harry-fruit/ddd-go/config"
	httpserver "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	port := fmt.Sprintf(":%s", config.ServerPort)

	server := httpserver.NewHTTPServer(port)
	server.Start()
}
