package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {

	ctx := context.Background()

	store := NewStore(testDB)

	user1 := createNewTestUser(t)
	user2 := createNewTestUser(t)

	// Note, using smaller, nice numbers for easier debugging
	account1Params := createRandomNewAccount(user1.Username)
	account1Params.Balance = 100
	account2Params := createRandomNewAccount(user2.Username)
	account2Params.Balance = 100

	account1, err := store.CreateAccount(ctx, account1Params)
	require.NoError(t, err, "error creating account 1")
	account2, err := store.CreateAccount(ctx, account2Params)
	require.NoError(t, err, "error creating account 2")

	// Run concurrent transfer txs

	n := 5

	// Connect the main go routine to the children with a channel
	errs := make(chan error)
	results := make(chan TransferTxResult)
	amount := int64(10)

	for i := 0; i < n; i++ {

		txName := fmt.Sprintf("tx %d", i)
		go func() {
			txCtx := context.WithValue(ctx, txKey, txName)
			result, err := store.TransferTx(txCtx, TransferTxParams{
				FromAccountId: account1.ID,
				ToAccountId:   account2.ID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	existedTxs := make(map[int]bool)
	// Check the results from each transaction in each go routine
	for i := 0; i < n; i++ {
		err := <-errs
		// Remember they may not be in the same order, I could label them though in a way
		require.NoError(t, err, fmt.Sprintf("transaction %d had an error", i))

		result := <-results
		require.NotEmpty(t, result, fmt.Sprintf("result %d is empty", i))

		require.NotEmpty(t, result.Transfer)
		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)
		require.Equal(t, amount, result.Transfer.Amount)
		require.NotEqual(t, "", result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)

		require.NotEmpty(t, result.FromAccount)
		require.Equal(t, account1.ID, result.FromAccount.ID)

		require.NotEmpty(t, result.ToAccount)
		require.Equal(t, account2.ID, result.ToAccount.ID)
		fromDiff := account1.Balance - result.FromAccount.Balance
		toDiff := result.ToAccount.Balance - account2.Balance
		// this is not a reasonable test because you can't guarantee the ordering
		// require.Equal(t, toDiff, amount * int64(i))
		// require.Equal(t, fromDiff, amount * int64(i))
		require.Equal(t, fromDiff, toDiff)
		require.True(t, fromDiff > 0)
		require.True(t, fromDiff%amount == 0)

		// Track which txs we have had by using the number of times the amount has gone into the diff so far
		k := int(fromDiff / amount)
		require.True(t, k >= 1 && k <= n)
		require.NotContains(t, existedTxs, k)
		existedTxs[k] = true

		require.NotEmpty(t, result.FromEntry)
		require.Equal(t, account1.ID, result.FromEntry.AccountID)
		require.Equal(t, -amount, result.FromEntry.Amount)
		require.NotEqual(t, "", result.FromEntry.ID)
		require.NotZero(t, result.FromEntry.CreatedAt)

		require.NotEmpty(t, result.ToEntry)
		require.Equal(t, account2.ID, result.ToEntry.AccountID)
		require.Equal(t, amount, result.ToEntry.Amount)
		require.NotEqual(t, "", result.ToEntry.ID)
		require.NotZero(t, result.ToEntry.CreatedAt)
	}

	totalChange := (int64(n) * amount)

	updatedFromAccount, err := store.GetAccount(ctx, account1.ID)
	require.NoError(t, err)
	require.Equal(t, account1.Balance-totalChange, updatedFromAccount.Balance)
	updatedToAccount, err := store.GetAccount(ctx, account2.ID)
	require.NoError(t, err)
	require.Equal(t, account2.Balance+totalChange, updatedToAccount.Balance)

}

// This test simulates a deadlock where concurrent tests are updating the same account rows in a different order at the same time
// So one (whichever goes first) will be given an exclusive lock on the row, and the other will attempt to receive a sharelock and ultimately cause a deadlock
// If the queries happen in a consistent order every time though, this will keep the deadlock from occuring.
// This is done by always using some "sorting" criteria for which is updated first, id comparison for instance
func TestTransferTxDeadlock(t *testing.T) {

	ctx := context.Background()

	store := NewStore(testDB)

	user1 := createNewTestUser(t)
	user2 := createNewTestUser(t)

	// Note, using smaller, nice numbers for easier debugging
	account1Params := createRandomNewAccount(user1.Username)
	account1Params.Balance = 100
	account2Params := createRandomNewAccount(user2.Username)
	account2Params.Balance = 100

	account1, err := store.CreateAccount(ctx, account1Params)
	require.NoError(t, err, "error creating account 1")
	account2, err := store.CreateAccount(ctx, account2Params)
	require.NoError(t, err, "error creating account 2")

	// Run concurrent transfer txs
	n := 10

	// Connect the main go routine to the children with a channel
	errs := make(chan error)
	amount := int64(10)

	for i := 0; i < n; i++ {

		txName := fmt.Sprintf("tx %d", i)

		// half send money from 1 -> 2
		// half send money from 2 -> 1
		fromAccountID := account1.ID
		toAccountID := account2.ID
		if i%2 == 1 {
			fromAccountID = account2.ID
			toAccountID = account1.ID
		}
		go func() {
			txCtx := context.WithValue(ctx, txKey, txName)
			_, err := store.TransferTx(txCtx, TransferTxParams{
				FromAccountId: fromAccountID,
				ToAccountId:   toAccountID,
				Amount:        amount,
			})

			errs <- err
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err, fmt.Sprintf("transaction %d had an error", i))
	}
	// Both account should end where they began because they both have an equal number of from and to transfers with the same money amount
	updatedAccount1, err := store.GetAccount(ctx, account1.ID)
	require.NoError(t, err)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	updatedAccount2, err := store.GetAccount(ctx, account2.ID)
	require.NoError(t, err)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}
