package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

func EncryptValue(key, text []byte) []byte {
	var block, cipherErr = aes.NewCipher(key)

	if cipherErr != nil {
		panic(cipherErr)
	}

	var aesgcm, gcmErr = cipher.NewGCM(block)

	if gcmErr != nil {
		panic(gcmErr)
	}

	var nonce = make([]byte, aesgcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}

	var ciphertext = aesgcm.Seal(nonce, nonce, text, nil)
	return ciphertext
}

func DecryptValue(key, data []byte) []byte {
	var block, cipherErr = aes.NewCipher(key)

	if cipherErr != nil {
		panic(cipherErr)
	}

	var aesgcm, gcmErr = cipher.NewGCM(block)

	if gcmErr != nil {
		panic(gcmErr)
	}

	var nonceSize = aesgcm.NonceSize()

	if len(data) < nonceSize {
		panic("Length of data is less than the nonce size, this can be due to incorrect encryption or corrupted data")
	}

	var nonce, cipherText = data[:nonceSize], data[nonceSize:]

	var text, decryptErr = aesgcm.Open(nil, nonce, cipherText, nil)

	if decryptErr != nil {
		panic(decryptErr)
	}

	return text
}
