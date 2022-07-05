package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyz"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func RandomString(n int) string {
	var sb strings.Builder

	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomMoney() int64 {
	// For the max, took the highest net worth currently and rounded up to the next order of magnitude higher
	return RandomInt(0, 1000000000000)
}

func RandomCurrency() string {
	currencies := [4]string{"jakata", "USD", "CAD", "EUR"}
	return currencies[RandomInt(0, int64(len(currencies))-1)]
}
