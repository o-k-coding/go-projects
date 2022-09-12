package gapi

import (
	"context"
	"database/sql"

	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/internal/password"
	"github.com/okeefem2/simple_bank/pb"
	"github.com/okeefem2/simple_bank/validate"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	err := maybeInvalidArgumentError(validateUpdateUserRequest(req))
	if err != nil {
		return nil, err
	}

	arg := db.UpdateUserParams{
		Username: req.GetUsername(),
		FullName: sql.NullString{
			String: req.GetFullName(),
			Valid:  req.FullName != nil,
		},
		Email: sql.NullString{
			String: req.GetEmail(),
			Valid:  req.Email != nil,
		},
	}

	if req.Password != nil {
		hashedPassword, err := password.HashPassword(req.GetPassword())
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to update user (hp) %s", err)
		}
		arg.HashedPassword = sql.NullString{
			String: hashedPassword,
			Valid:  true,
		}
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		// Ideally woould check for common errors like `unique_violation` etc and provide better error codes/messages
		return nil, status.Errorf(codes.Internal, "failed to update user (db) %s", err)
	}

	return &pb.UpdateUserResponse{
		User: convertUser(&user),
	}, nil
}

// Could update a struct that has a list of fields and functions to validate and a method to construct this
func validateUpdateUserRequest(req *pb.UpdateUserRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validate.ValidateUsername(req.GetUsername()); err != nil {
		violations = append(violations, fieldViolation("username", err))
	}

	if req.Password != nil {
		if err := validate.ValidatePassword(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("password", err))
		}
	}

	if req.FullName != nil {
		if err := validate.ValidateFullName(req.GetPassword()); err != nil {
			violations = append(violations, fieldViolation("fullName", err))
		}
	}

	if req.Email != nil {
		if err := validate.ValidateEmail(req.GetEmail()); err != nil {
			violations = append(violations, fieldViolation("emai;", err))
		}
	}

	return violations
}
