package utils

import (
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type hashParams struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	keyLength   uint32
}

func HashVaultPassword(password string) string {
	var params = &hashParams{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		keyLength:   128,
	}

	var passwordHash = generateHashFromPassword(password, params)
	var encodedPassword = base64.RawStdEncoding.EncodeToString(passwordHash)

	return encodedPassword
}

func generateHashFromPassword(password string, params *hashParams) []byte {
	var hash = argon2.IDKey(
		[]byte(password),
		[]byte{},
		params.iterations,
		params.memory,
		params.parallelism,
		params.keyLength)

	return hash
}
