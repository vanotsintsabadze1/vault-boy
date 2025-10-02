package misc

import (
	"os"
	"path/filepath"
)

func GetVaultsDir() string {
	home, err := os.UserConfigDir()

	if err != nil {
		panic(err)
	}

	return filepath.Join(home, AppDirName)
}

func GetDbNameWithExtension(value string) string {
	var combined = value + ".db"
	return combined
}
