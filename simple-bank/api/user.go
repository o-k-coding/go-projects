package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/okeefem2/simple_bank/db/sqlc"
	"github.com/okeefem2/simple_bank/internal/password"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=10"`
	FullName string `json:"fullName" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type userResponse struct {
	ID                uuid.UUID `json:"id"`
	Username          string    `json:"username"`
	FullName          string    `json:"fullName"`
	Email             string    `json:"email"`
	CreatedAt         time.Time `json:"createdAt"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
}

// This is done in order to make sure sensitive data is not returned
func newUserResponse(user db.User) userResponse {
	return userResponse{
		user.ID,
		user.Username,
		user.FullName,
		user.Email,
		user.CreatedAt,
		user.PasswordChangedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	// Get data from post body
	var body createUserRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := password.HashPassword(body.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}

	arg := db.CreateUserParams{
		Username:       body.Username,
		HashedPassword: hashedPassword,
		FullName:       body.FullName,
		Email:          body.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		handleApiError(ctx, err)
		return
	}

	// Use a response type so that the hashed password isn't returned as well.
	ctx.JSON(http.StatusCreated, newUserResponse(user))
}

type loginUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginUserResponse struct {
	AccessToken string `json:"accessToken"`
	User        userResponse
}

func (server *Server) loginUser(ctx *gin.Context) {
	var body loginUserRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, body.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.Status(http.StatusUnauthorized)
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	err = password.CheckPassword(user.HashedPassword, body.Password)
	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	accessToken, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}

	ctx.JSON(http.StatusOK, rsp)
}
