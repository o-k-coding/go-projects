package db

import (
	"context"
	"database/sql"
	"sort"
	"testing"

	"github.com/okeefem2/simple_bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomNewAccount(owner string) CreateAccountParams {
	return CreateAccountParams{
		Owner:    owner,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}

func createNewTestAccount(t *testing.T) Account {
	user := createNewTestUser(t)
	params := createRandomNewAccount(user.Username)

	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
	return account
}

func TestCreateAccount(t *testing.T) {
	user := createNewTestUser(t)
	params := createRandomNewAccount(user.Username)

	account, err := testQueries.CreateAccount(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, params.Owner, account.Owner)
	require.Equal(t, params.Balance, account.Balance)
	require.Equal(t, params.Currency, account.Currency)
	require.NotEqual(t, "", account.ID)
	require.NotZero(t, account.CreatedAt)
}

func TestGetAccount(t *testing.T) {
	ctx := context.Background()
	account := createNewTestAccount(t)

	fetchedAccount, err := testQueries.GetAccount(ctx, account.ID)
	require.NoError(t, err, "error getting new account")
	require.NotEmpty(t, fetchedAccount, "fetched account is empty")
	require.Equal(t, account, fetchedAccount, "created and fetched accounts not equal")
}

func TestDeleteAccount(t *testing.T) {
	ctx := context.Background()
	account := createNewTestAccount(t)

	err := testQueries.DeleteAccount(ctx, account.ID)
	require.NoError(t, err, "error deleting account")
	account, err = testQueries.GetAccount(ctx, account.ID)
	require.EqualError(t, err, sql.ErrNoRows.Error(), "wrong error getting deleted account")
	require.Empty(t, account, "deleted account is not empty")
}

func TestUpdateAccountBalance(t *testing.T) {
	ctx := context.Background()
	account := createNewTestAccount(t)

	updateParams := UpdateAccountBalanceParams{
		ID:      account.ID,
		Balance: account.Balance + 10,
	}

	updatedAccount, err := testQueries.UpdateAccountBalance(ctx, updateParams)
	require.NoError(t, err, "error updating account")
	require.NotEmpty(t, updatedAccount, "fetched account is empty")
	require.Equal(t, account.ID, updatedAccount.ID, "created and updated account ids are not equal")
	require.NotEqual(t, account.Balance, updatedAccount.Balance, "created and updated account balances are equal")
}

func TestListAccounts(t *testing.T) {
	ctx := context.Background()
	var lastAccount Account

	// Just to ensure there are least 6 entries in the DB to work with.
	for i := 0; i < 6; i++ {
		lastAccount = createNewTestAccount(t)
	}

	params := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Offset: 0,
		Limit:  5,
	}

	accounts, err := testQueries.ListAccounts(ctx, params)
	require.NoError(t, err, "error listing accounts")
	// require.Len(t, accounts, 5, "incorrect number of accounts listed")
	require.NotEmpty(t, accounts)

	// The accounts should be sorted by name
	sortedAccounts := make([]Account, len(accounts))
	copy(sortedAccounts, accounts)
	sort.Slice(sortedAccounts, func(i, j int) bool {
		// Wonder if string comp would be faster?
		// Note I had these ordered by name before, but name is not unique, so this causes some trouble
		// Nicely though, this test caught that sly bug. if it is a bug? idk not necessarily.
		return sortedAccounts[i].CreatedAt.Before(sortedAccounts[j].CreatedAt)
	})

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount, account)
	}
}
