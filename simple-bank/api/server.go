package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/okeefem2/simple_bank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()

	server := &Server{store, router}

	// Add routes

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
