package gapi

import (
	"context"
	"database/sql"

	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/internal/password"
	"github.com/okeefem2/simple_bank/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.Unauthenticated, "nah brah")
		}
		return nil, status.Errorf(codes.Internal, "failed to login %s", err)
	}
	err = password.CheckPassword(user.HashedPassword, req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "nah brah")
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating access token %s", err)
	}

	refreshToken, err := server.tokenMaker.CreateToken(
		user.Username,
		server.config.RefreshTokenDuration,
	)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating refresh token %s", err)
	}

	// In the course he added the payload as a return value of CreateToken, I am not going to because I am alittle lazy lol
	// I can see the benefit. I also could see it being a problem adjusting an existing API, maybe a security concern.
	// Unless this added significant performance issues, I would just leave it.
	refreshPayload, err := server.tokenMaker.VerifyToken(refreshToken)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating refresh payload %s", err)
	}

	m := server.extractMetadata(ctx)
	// NOTE: I don't think this is fully secure (secure code warrior, I think encoding/encryption is necessary?)
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		// These come from the gin context
		UserAgent: m.UserAgent,
		ClientIp:  m.ClientIP,
		IsBlocked: false,
		ExpiresAt: refreshPayload.ExpiredAt,
	})

	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error creating session %s", err)
	}

	rsp := pb.LoginUserResponse{
		SessionId:    session.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User:         convertUser(&user),
	}

	return &rsp, nil
}
