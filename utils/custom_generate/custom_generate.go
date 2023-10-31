package customgenerate

import (
	"math/rand"
	"time"
)

func GenerateTransactionCode() string {
	rand.Seed(time.Now().UnixNano())
	length := 16

	const availableChars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(availableChars))
		code[i] = availableChars[randomIndex]
	}

	return string(code)
}

func GenerateContractNumber() string {
	rand.Seed(time.Now().UnixNano())
	length := 8

	const availableChars = "0123456789"
	code := make([]byte, length)
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(availableChars))
		code[i] = availableChars[randomIndex]
	}
	return string(code)
}
