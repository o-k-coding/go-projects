package gapi

import (
	"context"

	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/internal/password"
	"github.com/okeefem2/simple_bank/pb"
	"github.com/okeefem2/simple_bank/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	err := maybeInvalidArgumentError(validateCreateUserRequest(req))
	if err != nil {
		return nil, err
	}
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

// Could create a struct that has a list of fields and functions to validate and a method to construct this
func validateCreateUserRequest(req *pb.CreateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if err := validate.ValidatePassword(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("password", err))
	}

	if err := validate.ValidateFullName(req.GetPassword()); err != nil {
		violations = append(violations, fieldViolation("fullName", err))
	}

	if err := validate.ValidateEmail(req.GetEmail()); err != nil {
		violations = append(violations, fieldViolation("emai;", err))
	}

	return violations
}
