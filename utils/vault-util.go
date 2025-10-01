package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/vanotsintsabadze1/vault-boy/misc"
)

func InitializeVaultDir() {
	home, configErr := os.UserConfigDir()

	if configErr != nil {
		panic(configErr)
	}

	var vault, password = getVaultAndPw()

	var vaultPath = filepath.Join(home, misc.AppDirName, vault)
	var err = os.MkdirAll(vaultPath, 0755)

	if err != nil {
		panic(err)
	}

	var hashedPw = HashVaultPassword(password)

	fmt.Println(hashedPw)
}

func getVaultAndPw() (string, string) {
	var vault, password string

	fmt.Print("Vault: ")
	fmt.Scan(&vault)

	fmt.Printf("Password for vault %s: ", vault)
	fmt.Scan(&password)

	return vault, password
}
