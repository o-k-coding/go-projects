package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/okeefem2/simple_bank/db/sqlc"
)

type createAccountRequest struct {
	Owner string `json:"owner" binding:"required"`
	// The oneof binding is definitely not very dynamic, would be interested if there are other options with gin
	// probably better to validate ourselves for anything more complex.
	Currency string `json:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	// Get data from post body
	var body createAccountRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    body.Owner,
		Currency: body.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, account)
}

type getAccountRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	// Get data from post body
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	accountID, err := uuid.Parse(req.ID)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccount(ctx, accountID)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	ctx.JSON(http.StatusOK, account)
}

// NOTE props must be public for tags to work properly lol. DINGUS
type listAccountsRequest struct {
	PageSize int32 `form:"pageSize" binding:"required,min=1"`
	Page     int32 `form:"page" binding:"required,min=1"`
}

type ListAccountsResponse struct {
	Accounts []db.Account `json:"accounts"`
	Page     int32        `json:"page"`
	PageSize int32        `json:"pageSize"`
	// Total int32 this one is a nice addition but requires some change to the db code
}

func (server *Server) listAccounts(ctx *gin.Context) {
	// Get data from post body
	var req listAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Bindings will always catch invalid state, so this is not needed
	// I am leaving this in as an example for me though
	req = setListPageDefaults(req)

	listParams := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: calculatePageOffset(req.PageSize, req.Page),
	}

	accounts, err := server.store.ListAccounts(ctx, listParams)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	res := ListAccountsResponse{
		accounts,
		req.Page,
		req.PageSize,
	}

	ctx.JSON(http.StatusOK, res)
}

func setListPageDefaults(req listAccountsRequest) listAccountsRequest {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}
	return req
}

func calculatePageOffset(pageSize int32, pageNum int32) int32 {
	return pageSize * (pageNum - 1)
}
