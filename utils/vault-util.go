package utils

import (
	"crypto/rand"
	"os"
	"path/filepath"

	"github.com/vanotsintsabadze1/vault-boy/misc"
)

func InitializeVaultDir() {
	vaultsDir := misc.GetVaultsDir()

	var vault, _ = GetVaultAndPw()
	var vaultPath = filepath.Join(vaultsDir, vault)

	createVaultDirectory(vaultPath)

	var vaultDbName = misc.GetDbNameWithExtension(vault)
	var vaultDbPath = filepath.Join(vaultPath, vaultDbName)

	db := CreateOrOpenVaultDb(vaultDbPath)

	// var derivedUsrPwKey = GenerateArgon2IDKey(password, 128)
	var masterKey = generateMasterKey()

	var bucket = CreateMetadataBucket(db)

	bucket.Put([]byte("master-key"), masterKey)

	defer db.Close()
}

func createVaultDirectory(vaultPath string) {
	var err = os.MkdirAll(vaultPath, 0755) // 0755 means read, write & execute

	if err != nil {
		panic(err)
	}
}

func generateMasterKey() []byte {
	masterKey := make([]byte, 32)
	rand.Read(masterKey)

	return masterKey
}
