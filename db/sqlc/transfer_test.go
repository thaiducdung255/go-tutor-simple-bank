package db

import (
	"testing"

	"github.com/thaiducdung255/simplebank/util"
)

func TestCreateTransfer(t *testing.T) {
	fromAccount := CreateRandomAccount(t)
	toAccount := CreateRandomAccount(t)

	arg := CreateTransferParams{
		FromAccountID: fromAccount.ID,
		ToAccountID:   toAccount.ID,
		Amount:        util.RandomInt(int64(0), fromAccount.Balance),
	}

	CreateAccountParams{}
}
