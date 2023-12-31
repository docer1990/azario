package main

import (
	"log"

	"github.com/docer1990/azario/api"
	"github.com/docer1990/azario/db"
	"github.com/docer1990/azario/utils"
)

func main() {
	utils.LoadEnvData(".env")
	store, err := db.NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3000", store)
	server.Run()
}
