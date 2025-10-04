package utils

import "go.etcd.io/bbolt"

func CreateOrOpenVaultDb(vaultName string) *bbolt.DB {
	db, err := bbolt.Open(vaultName, 0600, nil)

	if err != nil {
		panic(err)
	}

	return db
}

func CreateMetadataBucket(db *bbolt.DB) *bbolt.Bucket {
	var metaBucket *bbolt.Bucket
	err := db.Update(func(tx *bbolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("metadata"))

		if err != nil {
			tx.Rollback()
			return err
		}

		metaBucket = bucket
		return nil
	})

	if err != nil {
		panic(err)
	}

	return metaBucket
}
