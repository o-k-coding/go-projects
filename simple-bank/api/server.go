package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	customvalidator "github.com/okeefem2/simple_bank/api/custom_validator"
	db "github.com/okeefem2/simple_bank/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	router := gin.Default()

	server := &Server{store, router}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", customvalidator.ValidCurrency)
	}

	// Add routes

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	router.POST("/transfers", server.createTransfer)

	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
