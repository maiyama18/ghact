package cmd

import (
	"errors"
	"github.com/spf13/cobra"
)

var filename string
var token string

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Increase number of today's activity by 1",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("commit requires your repository name to commit")
		}
		return cobra.OnlyValidArgs(cmd, args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		return
	},
}

func init() {
	showCmd.Flags().StringVarP(&filename, "file", "f", ".ghact", "filename to be updated")
	showCmd.Flags().StringVarP(&token, "token", "t", "", "your private github token")
	_ = showCmd.MarkFlagRequired("token")
}
