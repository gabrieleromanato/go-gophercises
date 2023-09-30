package cmd

import (
	"fmt"
	"gabrieleromanato/secret/api"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a secret",
	Long:  `Get a secret`,
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
		if len(args) < 1 {
			fmt.Println("You must provide a key name")
			return
		}
		keyName := args[0]
		secret, err := api.ReadSecretFromFile(".secrets/" + keyName)
		if err != nil {
			fmt.Println("Error reading secret")
			return
		}
		decodedSecret, err := api.Decrypt(string(secret), key)
		if err != nil {
			fmt.Println("Error decrypting secret")
			return
		}
		fmt.Println(string(decodedSecret))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.Flags().StringP("key", "k", "passphrasewhichneedstobe32bytes!", "Secret key")
}
