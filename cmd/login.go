package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var username string
var token string

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to GitHub",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(username)
		fmt.Println(token)
	},
}

func init() {
	loginCmd.Flags().StringVarP(&username, "username", "u", "", "github username")
	loginCmd.Flags().StringVarP(&token, "token", "t", "", "github personal token")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("token")
}
