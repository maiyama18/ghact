package cmd

import (
	"errors"
	"fmt"
	"ghact/gh"
	"github.com/spf13/cobra"
)

var client = gh.NewClient()

var filepath string
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
		repo := args[0]
		file, err := fetchFile(repo, filepath)
		if err != nil {
			fmt.Printf("failed to fetch file from github: %s\n", err.Error())
			return
		}

		fmt.Println(file)

		return
	},
}

func init() {
	commitCmd.Flags().StringVarP(&filepath, "file", "f", ".ghact", "filepath to be updated")
	commitCmd.Flags().StringVarP(&token, "token", "t", "", "your private github token")
	_ = showCmd.MarkFlagRequired("token")
}

func fetchFile(repo, filepath string) (*GithubFile, error) {

}
