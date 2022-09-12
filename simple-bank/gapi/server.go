package gapi

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/okeefem2/simple_bank/config"
	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/pb"
	"github.com/okeefem2/simple_bank/token"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(store db.Store, config config.Config) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		// Should return an error from this function
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
