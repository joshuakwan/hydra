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

var createCmd = &cobra.Command{
	Use:   "create -f <configuration file>",
	Short: "create a resource from the config file",
	Run:   handleCreate,
}

var filename string

// Execute is the main logic of the CLI
func Execute() {
	createCmd.Flags().StringVarP(&filename, "file", "f", "", "source file to read from")
	createCmd.MarkFlagRequired("file")

	rootCmd.AddCommand(getCmd)
	rootCmd.AddCommand(createCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
