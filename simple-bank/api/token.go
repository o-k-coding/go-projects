package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken string `json:"accessToken"`
}

func (server *Server) renewAccessToken(ctx *gin.Context) {
	var body renewAccessTokenRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	refreshToken, err := server.tokenMaker.VerifyToken(body.RefreshToken)

	if err != nil {
		log.Println("error verifying token")
		ctx.Status(http.StatusUnauthorized)
		return
	}

	session, err := server.store.GetSession(ctx, refreshToken.ID)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no session found")
			log.Println(refreshToken.ID)
			ctx.Status(http.StatusUnauthorized)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	isExpired := time.Now().After(session.ExpiresAt)

	// Validate the session - should probably be another function
	if isExpired {
		log.Println("expired")
		ctx.Status(http.StatusUnauthorized)
		return
	}

	if session.Username != refreshToken.Username {
		log.Println("invalid username")
		ctx.Status(http.StatusUnauthorized)
		return
	}

	if session.RefreshToken != body.RefreshToken {
		log.Println("invalid refresh token")
		ctx.Status(http.StatusUnauthorized)
		return
	}

	if session.IsBlocked {
		log.Println("blocked")
		ctx.Status(http.StatusUnauthorized)
		return
	}

	// if isExpired || session.IsBlocked || session.Username != refreshToken.Username || session.RefreshToken != body.RefreshToken {
	// 	ctx.Status(http.StatusUnauthorized)
	// 	return
	// }

	// Maybe also check to make sure the user is actually good to go in a real system (not a locked user etc.)

	accessToken, err := server.tokenMaker.CreateToken(
		refreshToken.Username,
		server.config.AccessTokenDuration,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := renewAccessTokenResponse{
		AccessToken: accessToken,
	}

	ctx.JSON(http.StatusOK, rsp)
}
