package token

import (
	"time"

	"github.com/google/uuid"
)

// var ErrTokenExpired

type Payload struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	issuedAt := time.Now()

	return &Payload{
		ID:        tokenID,
		Username:  username,
		IssuedAt:  issuedAt,
		ExpiredAt: issuedAt.Add(duration),
	}, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrTokenExpired
	}
	return nil
}
