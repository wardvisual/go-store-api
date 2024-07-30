package main

import (
	"fmt"
	"log"

	"github.com/wardvisual/go-store-api/cmd/api"
)

func main() {

	server := api.NewAPIServer(fmt.Sprintf(":%s", 4000))

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
