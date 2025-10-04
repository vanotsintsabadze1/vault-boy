package utils

import (
	"fmt"
)

func GetVaultAndPw() (string, string) {
	var vault, password string

	fmt.Print("Vault: ")
	fmt.Scan(&vault)

	fmt.Printf("Password for vault %s: ", vault)
	fmt.Scan(&password)

	return vault, password
}
