package cmd

import (
	"fmt"
	"ghact/gh"
	"github.com/spf13/cobra"
	"os"
)

const shortDesc = "ghact is a CLI tool for viewing and manipulating your github activity"
const longDesc = `ghact is a CLI tool for viewing and manipulating your github activity.
documentation is available on https://github.com/muiscript/ghact`

var ghClient = gh.NewClient()

var rootCmd = &cobra.Command{
	Use:   "ghact",
	Short: shortDesc,
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello, ghact!")
	},
}

func init() {
	cobra.OnInitialize()

	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(loginCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
