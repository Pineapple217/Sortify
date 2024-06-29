package util

import (
	"math/rand"
	"time"
)

const charsetForState = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandomString(l int) string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, l)
	for i := range b {
		b[i] = charsetForState[seededRand.Intn(len(charsetForState))]
	}
	return string(b)
}
