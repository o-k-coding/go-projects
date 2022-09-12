package gapi

import (
	"context"

	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/internal/password"
	"github.com/okeefem2/simple_bank/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	// Better to use the Get functions because they add safety checks for nilness
	hashedPassword, err := password.HashPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user (hp) %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
		Email:          req.GetEmail(),
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		// Ideally woould check for common errors like `unique_violation` etc and provide better error codes/messages
		return nil, status.Errorf(codes.Internal, "failed to create user (db) %s", err)
	}

	return &pb.CreateUserResponse{
		User: convertUser(&user),
	}, nil
}
