package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrasnferTx(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	//run concurrent transfer transactions
	n := 2
	amount := int64(10)

	errs := make(chan error)
	//results := make(chan TransferTxResult)
	//existed := make(map[int]bool)
	for i := 0; i < n; i++ {

		fromAccountId := account1.ID
		toAccountId := account2.ID

		txName := fmt.Sprintf("tx %d", i+1)

		/*
			if i%2 == 1 {
				fromAccountId := account2.ID
				toAccountId := account1.ID
			}
		*/

		go func() {
			ctx := context.WithValue(context.Background(), txKey, txName)
			_, err := store.TransferTx(ctx, TransferTxParams{
				FromAccountID: fromAccountId,
				ToAccountID:   toAccountId,
				Amount:        amount,
			})

			errs <- err
			//results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)
	}

	// check the final updated balances

	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	//updatedAccount1Int, err := strconv.ParseInt(updatedAccount1.Balance, 10, 64)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	//updatedAccount2Int, err := strconv.ParseInt(updatedAccount2.Balance, 10, 64)
	require.NoError(t, err)

	require.Equal(t, updatedAccount1.Balance-int64(n)*amount, updatedAccount1.Balance)
	require.Equal(t, updatedAccount2.Balance+int64(n)*amount, updatedAccount2.Balance)
}
