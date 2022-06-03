package db

import (
	"context"
	"database/sql"
	"fmt"
	"sort"

	"github.com/google/uuid"
)

// Provide functions to exec db queries and txs
// This is a composition exposing the Queries functionality with the added Transaction functionality
// prefer composition over inheritance!
// I could see this as the public API for a package over the queries piece
type Store struct {
	*Queries
	db *sql.DB // Required to create new transaciton
}

func Newstore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// only expose functions that use this to do some unit of work
func (store *Store) execTransaction(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx error: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountId uuid.UUID `json:"fromAccountId"`
	ToAccountId uuid.UUID `json:"toAccountId"`
	Amount int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"fromAccount"`
	ToAccount Account `json:"toAccount"`
	FromEntry Entry `json:"fromEntry"`
	ToEntry Entry `json:"toEntry"`
}

var txKey = struct{}{}

// Perform a money transfer from one account to another
func (store *Store) TransferTx(ctx context.Context, params TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTransaction(ctx, func(q *Queries) error {
		var err error

		// txName := ctx.Value(txKey)

		// fmt.Println(txName, "create transfer")
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			params.FromAccountId,
			params.ToAccountId,
			params.Amount,
		})

		if err != nil {
			return err
		}

		// fmt.Println(txName, "create from entry")
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: params.FromAccountId,
			Amount: -params.Amount,
		})

		if err != nil {
			return err
		}

		// fmt.Println(txName, "create to entry")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: params.ToAccountId,
			Amount: params.Amount,
		})

		if err != nil {
			return err
		}

		// This is the first implemenation, not used in favor of the more efficient single update query
		// fmt.Println(txName, "get from account")
		// fromAccount, err := q.GetAccountForUpdate(ctx, params.FromAccountId)

		// if err != nil { return err }

		// // result.FromAccount = fromAccount
		// fmt.Println(txName, "udpdate from account")
		// result.FromAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		// 	ID: params.FromAccountId,
		// 	Balance: fromAccount.Balance - params.Amount,
		// })

		// if err != nil { return err }

		// fmt.Println(txName, "get to account")
		// toAccount, err := q.GetAccountForUpdate(ctx, params.ToAccountId)

		// if err != nil { return err }

		// // result.ToAccount = toAccount
		// fmt.Println(txName, "udpdate to account")
		// result.ToAccount, err = q.UpdateAccountBalance(ctx, UpdateAccountBalanceParams{
		// 	ID: params.ToAccountId,
		// 	Balance: toAccount.Balance + params.Amount,
		// })

		// if err != nil { return err }

		// If multiple concurrent txns involving the same account in different "orientation", we could still have a deadlock
		// So we solve it by always executing in a consistent order based on the account id
		addBalanceParams := []AddAccountBalanceParams{
			{
				ID: params.FromAccountId,
				Amount: -params.Amount,
			},
			{
				ID: params.ToAccountId,
				Amount: params.Amount,
			},
		}

		sort.Slice(addBalanceParams, func (i, j int) bool  {
			return addBalanceParams[i].ID.String() < addBalanceParams[j].ID.String()
		})

		for _, param := range addBalanceParams {
			account, err := q.AddAccountBalance(ctx, param)
			if err != nil {
				return err
				// Note that if for some reason the accounts are the same id
				// then you will never get the ToAccount here
				// I would either not make it possible to get this far if they are the same
				// or I would check the existance of the account on the result, and the amount as well
			} else if (param.ID == params.FromAccountId) {
				result.FromAccount = account
			} else {
				result.ToAccount = account
			}
		}

		if err != nil { return err }

		return nil
	})

	return result, err
}
