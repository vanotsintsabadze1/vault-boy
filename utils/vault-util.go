package utils

import (
	"os"
	"path/filepath"

	"github.com/vanotsintsabadze1/vault-boy/misc"
)

func InitializeVaultDir() {
	vaultsDir := misc.GetVaultsDir()

	var vault, password = GetVaultAndPw()
	var vaultPath = filepath.Join(vaultsDir, vault)

	createVaultDirectory(vaultPath)

	var vaultDbName = misc.GetDbNameWithExtension(vault)
	var vaultDbPath = filepath.Join(vaultPath, vaultDbName)

	db := CreateOrOpenVaultDb(vaultDbPath)

	HashVaultPassword(password)

	defer db.Close()
}

func createVaultDirectory(vaultPath string) {
	var err = os.MkdirAll(vaultPath, 0755) // 0755 means read, write & execute

	if err != nil {
		panic(err)
	}
}
