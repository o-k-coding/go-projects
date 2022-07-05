package password

import (
	"testing"

	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPassword(t *testing.T) {
	password := util.RandomString(6)

	hash, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = CheckPassword(hash, password)

	require.NoError(t, err)

	err = CheckPassword(hash, util.RandomString(6))

	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestPasswordWrong(t *testing.T) {
	password := util.RandomString(6)

	hash, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hash)

	err = CheckPassword(hash, util.RandomString(6))

	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}

func TestPasswordMultipleHash(t *testing.T) {
	password := util.RandomString(6)

	hash1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash1)
	hash2, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEmpty(t, hash2)

	require.NotEqual(t, hash1, hash2)
}
