package main

import (
	"log"

	_ "github.com/golang/mock/mockgen/model"
	_ "github.com/lib/pq"
	"github.com/okeefem2/simple_bank/api"
	"github.com/okeefem2/simple_bank/config"
	db "github.com/okeefem2/simple_bank/db/sqlc"
)

func main() {
	// SO a note here, this pattern is more about passing objects needed around
	// the other was more about creating objects that had access to the things needed,
	// then creating receiver functions on those. A more OOP approach.
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	conn := db.ConnectPostgres(config)
	store := db.NewStore(conn)
	server, err := api.NewServer(store, *config)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
