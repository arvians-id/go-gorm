package main

import (
	"github.com/arvians-id/go-gorm/cmd/config"
	"github.com/arvians-id/go-gorm/injection"
	"log"
	"net/http"
)

func main() {
	configuration := config.New()
	routes, err := injection.InitServerAPI(configuration)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(":8080", routes)
	if err != nil {
		log.Fatal("failed to start server")
	}
}
