package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/okeefem2/simple_bank/internal/password"
	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomNewUser() CreateUserParams {
	hashedPassword, _ := password.HashPassword(util.RandomString(6))
	return CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: hashedPassword,
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
}

func createNewTestUser(t *testing.T) User {
	params := createRandomNewUser()

	user, err := testQueries.CreateUser(context.Background(), params)
	require.NoError(t, err)
	return user
}

func TestCreateUser(t *testing.T) {
	params := createRandomNewUser()

	user, err := testQueries.CreateUser(context.Background(), params)
	require.NoError(t, err)

	require.NotEmpty(t, user)

	require.Equal(t, params.Username, user.Username)
	require.Equal(t, params.HashedPassword, user.HashedPassword)
	require.Equal(t, params.FullName, user.FullName)
	require.Equal(t, params.Email, user.Email)
	require.NotEqual(t, "", user.ID)
	require.NotZero(t, user.CreatedAt)
}

func TestGetUser(t *testing.T) {
	ctx := context.Background()
	user := createNewTestUser(t)

	fetchedUser, err := testQueries.GetUser(ctx, user.Username)
	require.NoError(t, err, "error getting new user")
	require.NotEmpty(t, fetchedUser, "fetched user is empty")
	require.Equal(t, user, fetchedUser, "created and fetched users not equal")
}

func TestUpdateUserOnlyFullName(t *testing.T) {
	user := createNewTestUser(t)

	newFullName := util.RandomOwner()
	updatedUser, err := testQueries.UpdateUser(context.Background(), UpdateUserParams{
		Username: user.Username,
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
	})

	require.NoError(t, err, "error getting new user")
	require.NotEmpty(t, updatedUser, "updated user is empty")
	require.Equal(t, user.Username, updatedUser.Username, "created and updated user username not equal")
	require.Equal(t, user.HashedPassword, updatedUser.HashedPassword, "created and updated user hashed password not equal")
	require.Equal(t, user.Email, updatedUser.Email, "created and updated user email not equal")
	require.Equal(t, newFullName, updatedUser.FullName, "updated user full name incorrect")
}
