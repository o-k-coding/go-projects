package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/okeefem2/simple_bank/db/sqlc"
)

type transferRequest struct {
	FromAccountId string `json:"fromAccountId" binding:"required,uuid"`
	ToAccountId   string `json:"toAccountId" binding:"required,uuid"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	// Get data from post body
	var body transferRequest

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// use parse, because must parse will cause a panic if it fails, we would rather handle the error ourselves
	fromAccountId, err := uuid.Parse(body.FromAccountId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	toAccountId, err := uuid.Parse(body.ToAccountId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateTransferParams{
		FromAccountID: fromAccountId,
		ToAccountID:   toAccountId,
		Amount:        body.Amount,
	}

	_, err = server.validAccount(ctx, arg.ToAccountID, body.Currency)

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		} else if _, ok := err.(*InvalidCurrencyError); ok {
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		}
		return
	}

	tx, err := server.store.CreateTransfer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, tx)
}

// This stuff should probably be in a file.
// Also never sure if we really want to make this a receiver function on server. Would rather business rules live at a different layer than the server.
type InvalidCurrencyError struct {
	toAccountId      uuid.UUID
	accountCurrency  string
	transferCurrency string
}

func (e *InvalidCurrencyError) Error() string {
	return fmt.Sprintf("account [%s] currency mismatch: %s vs %s", e.toAccountId.String(), e.accountCurrency, e.transferCurrency)
}

func (server *Server) validAccount(ctx *gin.Context, toAccountId uuid.UUID, currency string) (bool, error) {
	account, err := server.store.GetAccount(ctx, toAccountId)
	if err != nil {
		return false, err
	}

	if account.Currency != currency {
		return false, &InvalidCurrencyError{account.ID, account.Currency, currency}
	}

	return true, nil
}
