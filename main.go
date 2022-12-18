package main

import (
	"github.com/arvians-id/go-gorm/cmd/config"
	"github.com/arvians-id/go-gorm/cmd/server"
	"net/http"
)

func main() {
	configuration := config.New()
	chiMux := server.NewInitializedServer(configuration)

	err := http.ListenAndServe(":3000", chiMux)
	if err != nil {
		panic("failed to start server")
	}
}
