package cmd

import (
	"fmt"
	"ghact/config"
	"ghact/gh"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"path"
)

const shortDesc = "ghact is a CLI tool for viewing and manipulating your github activity"
const longDesc = `ghact is a CLI tool for viewing and manipulating your github activity.
documentation is available on https://github.com/muiscript/ghact`

var ghClient *gh.Client
var conf *config.Config

var rootCmd = &cobra.Command{
	Use:   "ghact",
	Short: shortDesc,
	Long:  longDesc,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello, ghact!")
	},
}

func init() {
	cobra.OnInitialize(func() {
		usr, err := user.Current()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configJsonPath := path.Join(usr.HomeDir, ".ghact.json")

		if conf, err = config.New(configJsonPath); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ghClient = gh.NewClient()
	})

	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(loginCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
