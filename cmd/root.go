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

func exit(code int, format string, a ...interface{}) {
	fmt.Printf(format, a)
	fmt.Println()

	os.Exit(code)
}

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
			exit(1, err.Error())
		}
		configJsonPath := path.Join(usr.HomeDir, ".ghact.json")

		if conf, err = config.New(configJsonPath); err != nil {
			exit(1, err.Error())
		}
		ghClient = gh.NewClient()
	})

	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(configCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		exit(1, err.Error())
	}
}
