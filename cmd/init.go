package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vanotsintsabadze1/vault-boy/utils"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"initialize"},
	Short:   "Initializes a vault to store passwords",
	Run: func(cmd *cobra.Command, args []string) {
		utils.InitializeVaultDir()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
