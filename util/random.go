package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt returns a random int64 between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n
func RandomString(n int) string {
	const ALPHABET = "qwertyuiopasdfghjklzxcvbnm"
	var sb strings.Builder
	k := len(ALPHABET)

	for i := 0; i < n; i++ {
		c := ALPHABET[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner returns a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney returns a random amount of money
func RandomMoney() int64 {
	return RandomInt(10, 5000)
}

var Currencies = []string{"USD", "CAD", "BELI", "YEN", "VND"}

// RandomCurrency returns a random unit of currency
func RandomCurrency() string {
	return Currencies[rand.Intn(len(Currencies))]
}
