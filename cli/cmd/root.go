package cmd

import (
	"os"

	"github.com/oleoneto/go-toolkit/logger"
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
			logg := logger.NewLogger(logger.LoggerOptions{Format: format})
			logg.Log(&version, os.Stdout, template)
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
	rootCmd.PersistentFlags().StringVarP(&format, "output-format", "o", format, "output format")
	rootCmd.PersistentFlags().StringVarP(&template, "output-template", "y", template, "template (used when output format is 'gotemplate')")

	// Commands
	rootCmd.AddCommand(versionCmd)
}
