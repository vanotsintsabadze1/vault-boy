package utils

import (
	"golang.org/x/crypto/argon2"
)

type hashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
}

func GenerateArgon2IDKey(password string, keyLength uint32) []byte {
	var params = &hashParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		keyLength:   keyLength,
	}

	var hash = argon2.IDKey(
		[]byte(password),
		[]byte{},
		params.iterations,
		params.memory,
		params.parallelism,
		params.keyLength)

	return hash
}
