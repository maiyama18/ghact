package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var username string
var token string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		if err := ghClient.Login(username, token); err != nil {
			fmt.Printf("failed to login: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("successfully logged in with user: %s\n", username)
	},
}

func init() {
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "github username")
	loginCmd.Flags().StringVarP(&token, "token", "t", "", "github personal token")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("token")
}
