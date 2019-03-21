package cmd

import (
	"errors"
	"fmt"
	"ghact/gh"
	"github.com/spf13/cobra"
	"time"
)

var filePath string

var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Increase number of today's activity by 1",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("commit requires your repository name to commit")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		username, err := conf.Get("username")
		if err != nil {
			exit(1, err.Error())
		}
		token, err := conf.Get("token")
		if err != nil {
			exit(1, err.Error())
		}

		message := fmt.Sprintf("commited by ghact: %v", time.Now())
		repo := args[0]
		file, err := ghClient.Fetch(username, repo, filePath)
		var commit *gh.Commit
		if err != nil {
			commit = gh.NewCommit("", message+"\n", message)
		} else {
			commit = gh.NewCommit(file.Sha, file.Content+message+"\n", message)
		}

		if err := ghClient.Create(username, repo, filePath, token, commit); err != nil {
			exit(1, err.Error())
		}
		fmt.Printf("successfully commited to %s/%s/%s\n", username, repo, filePath)
	},
}

func init() {
	commitCmd.Flags().StringVarP(&filePath, "file", "f", ".ghact", "filePath to be updated")
}
