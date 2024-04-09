package helpers

import (
	"crypto/rand"
	"math/big"
)

func GenerateRandString(num int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if num < 0 {
		num *= -1
	}

	randByteArray := make([]byte, num)

	for i := range randByteArray {
		randIndex, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		randByteArray[i] = charset[int64(randIndex.Int64())]
	}

	return string(randByteArray)
}
