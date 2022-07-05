package db

import (
	"context"
	"testing"

	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateTransfer(t *testing.T) {
	ctx := context.Background()

	accountOne := createNewTestAccount(t)
	accountTwo := createNewTestAccount(t)

	transferParams := CreateTransferParams{
		FromAccountID: accountOne.ID,
		ToAccountID:   accountTwo.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(ctx, transferParams)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transferParams.FromAccountID, transfer.FromAccountID)
	require.Equal(t, transferParams.ToAccountID, transfer.ToAccountID)
	require.Equal(t, transferParams.Amount, transfer.Amount)
	require.NotEqual(t, "", transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
}

func TestGetTransfer(t *testing.T) {
	ctx := context.Background()
	accountOne := createNewTestAccount(t)
	accountTwo := createNewTestAccount(t)

	transferParams := CreateTransferParams{
		FromAccountID: accountOne.ID,
		ToAccountID:   accountTwo.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(ctx, transferParams)
	require.NoError(t, err, "error creating new transfer")

	fetchedTransfer, err := testQueries.GetTransfer(ctx, transfer.ID)
	require.NoError(t, err, "error getting new transfer")
	require.NotEmpty(t, fetchedTransfer, "fetched transfer is empty")
	require.Equal(t, transfer, fetchedTransfer, "created and fetched transfers not equal")
}
