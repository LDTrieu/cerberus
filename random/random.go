package random

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const numset = "0123456789"

// nolint
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// StringRandom generate a random string
func StringRandom(length int) string {
	return stringWithCharset(length, charset)
}

// AlphaNumericRandom ...
func AlphaNumericRandom(length int) string {
	return stringWithCharset(length, charset+numset)
}

// NumericRandom ...
func NumericRandom(length int) string {
	return stringWithCharset(length, numset)
}

func HashRandFixed(s string) uint32 {
	var (
		seed  uint32 = 5381
		prime uint32 = 16777619
		value uint32 = seed
	)

	for _, c := range s {
		value = (value * prime) + uint32(c)
	}

	return value
}
