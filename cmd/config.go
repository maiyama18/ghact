package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show config",
	Run: func(cmd *cobra.Command, args []string) {
		keys := conf.Keys()
		for _, key := range keys {
			val, err := conf.Get(key)
			if err != nil {
				continue
			}

			fmt.Printf("%s: %s\n", key, val)
		}
	},
}

func init() {}
