package utils

import "go.etcd.io/bbolt"

func CreateOrOpenVaultDb(vaultName string) *bbolt.DB {
	db, err := bbolt.Open(vaultName, 0600, nil)

	if err != nil {
		panic(err)
	}

	return db
}
