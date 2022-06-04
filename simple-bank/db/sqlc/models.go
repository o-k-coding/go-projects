// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID        uuid.UUID `json:"id"`
	Owner     string    `json:"owner"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"createdAt"`
}

type Entry struct {
	ID        uuid.UUID `json:"id"`
	AccountID uuid.UUID `json:"accountID"`
	// Can be negative or positive
	Amount    int64     `json:"amount"`
	CreatedAt time.Time `json:"createdAt"`
}

type Transfer struct {
	ID            uuid.UUID `json:"id"`
	FromAccountID uuid.UUID `json:"fromAccountID"`
	ToAccountID   uuid.UUID `json:"toAccountID"`
	// must be positive
	Amount    int64        `json:"amount"`
	CreatedAt sql.NullTime `json:"createdAt"`
}