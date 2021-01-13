package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thaiducdung255/simplebank/util"
)

func createRandomTransfer(t *testing.T) Transfer {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)
	amount := util.RandomInt(int64(0), fromAccount.Balance)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        amount,
	}

	createdTransfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, createdTransfer.ID)
	require.NotNil(t, createdTransfer.CreatedAt)
	require.Equal(t, createdTransfer.FromAccountID, fromAccount.ID)
	require.Equal(t, createdTransfer.ToAccountID, toAccount.ID)
	require.Equal(t, createdTransfer.Amount, amount)

	return createdTransfer
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	gotTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotNil(t, gotTransfer)
	require.NotNil(t, gotTransfer.ID)
	require.NotNil(t, gotTransfer.CreatedAt)
	require.Equal(t, gotTransfer.ID, transfer.ID)
	require.Equal(t, gotTransfer.CreatedAt, transfer.CreatedAt)
	require.Equal(t, gotTransfer.FromAccountID, transfer.FromAccountID)
	require.Equal(t, gotTransfer.ToAccountID, transfer.ToAccountID)
	require.Equal(t, gotTransfer.Amount, transfer.Amount)
}

func TestDeleteTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	err := testQueries.DeleteTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	nonExistTransfer, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.Error(t, err)
	require.Empty(t, nonExistTransfer)
}

func TestListTransfer(t *testing.T) {
	n := util.RandomInt(0, 20)
	var createdTransfers []Transfer

	for i := 0; i < int(n); i++ {
		createdTransfers = append(createdTransfers, createRandomTransfer(t))
	}

	offset := int32(util.RandomInt(0, n))
	arg := ListTransfersParams{
		Offset: offset,
		Limit:  int32(n) - offset,
	}

	queriedTransfers, err := testQueries.ListTransfers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, queriedTransfers)
	require.Equal(t, int32(len(queriedTransfers)), arg.Limit)
}

func TestUpdateTransfer(t *testing.T) {
	transfer := createRandomTransfer(t)

	arg := UpdateTransferParams{
		ID:     transfer.ID,
		Amount: util.RandomMoney(),
	}

	updatedTransfer, err := testQueries.UpdateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, updatedTransfer)
	require.Equal(t, updatedTransfer.Amount, arg.Amount)
	require.Equal(t, updatedTransfer.ID, arg.ID)
}
