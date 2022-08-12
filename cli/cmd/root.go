package cmd

import (
	"os"

	"github.com/cleopatrio/cli/logger"
	"github.com/spf13/cobra"
)

var (
	config   string
	format   = "plain"
	template = ""

	rootCmd = &cobra.Command{
		Use:   "cli",
		Short: "CLI, a cli tool.",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Shows the version of the CLI",
		Run: func(cmd *cobra.Command, args []string) {
			logger.Custom(format, template).WithFormattedOutput(&version, os.Stdout)
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func initConfig() {}

func init() {
	// CLI configuration
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&config, "config", config, "config file")

	// Commands
	rootCmd.AddCommand(versionCmd)
}
