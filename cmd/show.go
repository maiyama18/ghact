package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use: "show",
	Short: "Show number of activities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is show command")
	},
}
