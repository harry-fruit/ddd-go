package main

import (
	"log"

	httpserver "github.com/harry-fruit/ddd-go/internal/infrastructure/api/http"
	di "github.com/harry-fruit/ddd-go/pkg/container_di"
)

func main() {
	container := di.BuildContainer()

	err := container.Invoke(func(api *httpserver.HTTPServer) {
		api.Start(container)
	})

	if err != nil {
		log.Fatalf("failed to invoke: %v", err)
	}

}
