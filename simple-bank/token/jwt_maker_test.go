package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)
	username := util.RandomOwner()
	duration := time.Minute
	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)
	token, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)

	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotZero(t, payload.ID)
	require.Equal(t, username, payload.Username)
	// Could make it less than a second too, maybe 100ms.
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, time.Second)
	require.Equal(t, payload.IssuedAt.Add(duration), payload.ExpiredAt, time.Second)
}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	token, err := maker.CreateToken(util.RandomOwner(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Nil(t, payload)
	require.Error(t, err, ErrTokenExpired.Error())
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(util.RandomOwner(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Nil(t, payload)
	require.Error(t, err, ErrInvalidToken.Error())
}
