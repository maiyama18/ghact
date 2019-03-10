package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
)

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show number of activities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is show command")
	},
}

func fetchActivityHTML(user string) {
	http.Get()
}
