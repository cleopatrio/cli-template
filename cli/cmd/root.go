package cmd

import (
	"github.com/cleopatrio/cli/cli/printer"
	"github.com/drewstinnett/gout/v2"
	"github.com/spf13/cobra"
)

var Printer *gout.Gout

// Flags
var (
	config string

	outputTemplate string

	outputFormat = &printer.OutputFormat{
		Options: []string{"plain", "json", "yaml", "table", "gotemplate"},
		Default: "plain",
	}
)

var (
	rootCmd = &cobra.Command{
		Use:   "cli",
		Short: "CLI, a cli tool.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			printer.SetFormatter(Printer, outputFormat.String(), outputTemplate)
		},
		Run: func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
)

func initConfig() {}

func Execute() error { return rootCmd.Execute() }

func init() {
	Printer = gout.New()

	// CLI configuration
	cobra.OnInitialize(initConfig)

	// Flags
	rootCmd.PersistentFlags().StringVar(&config, "config", config, "config file")
	rootCmd.PersistentFlags().VarP(outputFormat, "output", "o", "output format")
	rootCmd.PersistentFlags().StringVarP(&outputTemplate, "output-template", "y", outputTemplate, "template (used when output format is 'gotemplate')")

	// Commands
	rootCmd.AddCommand(versionCmd)
}
