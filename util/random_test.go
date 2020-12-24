package util

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRandomInt(t *testing.T) {
	const MIN int64 = 20
	const MAX int64 = 200
	var randomInt int64 = RandomInt(MIN, MAX)
	require.NotNil(t, randomInt)
	require.LessOrEqual(t, randomInt, MAX)
	require.GreaterOrEqual(t, randomInt, MIN)
}

func TestRandomString(t *testing.T) {
	const STR_lEN int = 21
	var randomStr string = RandomString(STR_lEN)
	require.NotNil(t, randomStr)
	require.Equal(t, len(randomStr), STR_lEN)
}

func TestRandomOwner(t *testing.T) {
	var randOwner string = RandomOwner()
	require.NotNil(t, randOwner)
	require.Equal(t, len(randOwner), 6)
}

func TestRandomMoney(t *testing.T) {
	var randMoney int64 = RandomMoney()
	require.NotNil(t, randMoney)
	require.LessOrEqual(t, randMoney, int64(5000))
	require.GreaterOrEqual(t, randMoney, int64(10))
}

func TestRandomCurrency(t *testing.T) {
	var randCurrency string = RandomCurrency()
	require.NotNil(t, randCurrency)

	var validCurrency bool = false

	for _, v := range Currencies {
		if v == randCurrency {
			validCurrency = true
			break
		}
	}

	require.Equal(t, validCurrency, true)
}
