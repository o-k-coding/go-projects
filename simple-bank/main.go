package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/okeefem2/simple_bank/api"
	db "github.com/okeefem2/simple_bank/db/sqlc"
)

func main() {
	conn := db.ConnectPostgres()
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err := server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
