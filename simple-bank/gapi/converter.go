package gapi

import (
	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(dbUser *db.User) *pb.User {
	return &pb.User{
		Username:          dbUser.Username,
		FullName:          dbUser.FullName,
		Email:             dbUser.Email,
		PasswordChangedAt: timestamppb.New(dbUser.PasswordChangedAt),
		CreatedAt:         timestamppb.New(dbUser.CreatedAt),
	}
}
