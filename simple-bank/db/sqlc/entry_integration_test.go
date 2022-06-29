package db

import (
	"context"
	"testing"

	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
)


func TestCreateEntry(t *testing.T) {
	ctx := context.Background()
	params := createRandomNewAccount()

	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)

	entryParams := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(ctx, entryParams)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entryParams.AccountID, entry.AccountID)
	require.Equal(t, account.ID, entry.AccountID)
	require.Equal(t, entryParams.Amount, entry.Amount)
	require.NotEqual(t, "", entry.ID)
	require.NotZero(t, entry.CreatedAt)
}

func TestGetEntry(t *testing.T) {
	ctx := context.Background()
	params := createRandomNewAccount()

	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)

	entryParams := CreateEntryParams{
		AccountID: account.ID,
		Amount: util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(ctx, entryParams)
	require.NoError(t, err, "error creating entry to get")

	fetchedEntry, err := testQueries.GetEntry(ctx, entry.ID)
	require.NoError(t, err, "error getting new entry")
	require.NotEmpty(t, fetchedEntry, "fetched entry is empty")
	require.Equal(t, entry, fetchedEntry, "created and fetched entries not equal")
}
