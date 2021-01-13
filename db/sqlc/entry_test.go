package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thaiducdung255/simplebank/util"
)

func createRandomEntry(t *testing.T) Entry {
	randomAccount := CreateRandomAccount(t)

	arg := CreateEntryParams{
		Amount:    util.RandomInt(0, randomAccount.Balance),
		AccountID: randomAccount.ID,
	}

	createdEntry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, createdEntry)
	require.NotNil(t, createdEntry.ID)
	require.NotNil(t, createdEntry.CreatedAt)
	require.Equal(t, createdEntry.AccountID, arg.AccountID)
	require.Equal(t, createdEntry.Amount, arg.Amount)

	return createdEntry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	createdEntry := createRandomEntry(t)

	gotEntry, err := testQueries.GetEntry(context.Background(), createdEntry.ID)
	require.NoError(t, err)
	require.NotNil(t, gotEntry)
	require.Equal(t, createdEntry.ID, gotEntry.ID)
	require.Equal(t, createdEntry.Amount, gotEntry.Amount)
	require.Equal(t, createdEntry.CreatedAt, gotEntry.CreatedAt)
	require.Equal(t, createdEntry.AccountID, gotEntry.AccountID)
}

func TestListEntries(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomEntry(t)
	}

	arg := ListEntriesParams{
		Limit:  5,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, entries)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}

func TestDeleteEntry(t *testing.T) {
	randomEntry := createRandomEntry(t)

	err := testQueries.DeleteEntry(context.Background(), randomEntry.ID)
	require.NoError(t, err)

	nonExistEntry, err := testQueries.GetEntry(context.Background(), randomEntry.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, nonExistEntry)
}

func TestUpdateEntry(t *testing.T) {
	randomEntry := createRandomEntry(t)

	newAccID := util.RandomInt(1, 500)

	arg := UpdateEntryParams{
		AccountID: newAccID,
		ID:        randomEntry.ID,
	}

	updatedEntry, err := testQueries.UpdateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, updatedEntry)
	require.Equal(t, updatedEntry.ID, randomEntry.ID)
	require.Equal(t, updatedEntry.CreatedAt, randomEntry.CreatedAt)
	require.Equal(t, updatedEntry.Amount, randomEntry.Amount)
	require.Equal(t, updatedEntry.AccountID, newAccID)
}
