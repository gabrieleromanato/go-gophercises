package cmd

import (
	"fmt"
	"gabrieleromanato/secret/api"

	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a secret",
	Long:  `Set a secret`,
	Run: func(cmd *cobra.Command, args []string) {
		key, err := cmd.Flags().GetString("key")
		if err != nil {
			fmt.Println("Invalid key")
			return
		}
		if len(key) != 32 {
			fmt.Println("Key must be 32 characters long")
			return
		}
		if len(args) < 2 {
			fmt.Println("You must provide a key name and a secret")
			return
		}
		keyName := args[0]
		secret := args[1]

		dirPath := api.CreateSecretDirIfNotExists(".secrets")
		filePath := dirPath + "/" + keyName
		if api.FileExists(filePath) {
			fmt.Println("Key already exists")
			return
		}
		encodedSecret, err := api.Encrypt(secret, key)
		if err != nil {
			fmt.Println("Error encrypting secret")
			return
		}
		if !api.WriteSecretToFile(encodedSecret, filePath) {
			fmt.Println("Error writing secret to file")
			return
		}
		fmt.Println("Secret set")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().StringP("key", "k", "passphrasewhichneedstobe32bytes!", "Secret key")
}
