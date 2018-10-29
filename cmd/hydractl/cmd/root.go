package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hydractl",
	Short: "CLI for Hydra",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var getCmd = &cobra.Command{
	Use:       "get [event/rules/actions]",
	Short:     "get a resource",
	ValidArgs: []string{"events", "ev", "rules", "ru", "actions", "ac"},
	Args:      cobra.OnlyValidArgs,
	Run:       handleGet,
}

func Execute() {
	rootCmd.AddCommand(getCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
