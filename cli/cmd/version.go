package cmd

import (
	"fmt"

	"github.com/cleopatrio/cli/cli/printer"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Shows the version of the CLI",
	Run:   func(cmd *cobra.Command, args []string) { Printer.MustPrint(version) },
}

type Version struct {
	Major string `json:"major" yaml:"major" toml:"major" xml:"major"`
	Minor string `json:"minor" yaml:"minor" toml:"minor" xml:"minor"`
	Patch string `json:"patch" yaml:"patch" toml:"patch" xml:"patch"`
}

func (V *Version) String() string { return fmt.Sprintf(`cli %v.%v.%v`, V.Major, V.Minor, V.Patch) }

func (v Version) TableWriter() table.Writer {
	title := "version"
	t := printer.Initialize(printer.TableOptions{
		Title:        &title,
		OutputMirror: nil,
		Style:        &table.StyleLight,
	})

	t.AppendHeader(table.Row{"major", "minor", "patch"})
	t.AppendRow(table.Row{v.Major, v.Minor, v.Patch})

	return t
}

var version = Version{Major: "0", Minor: "1", Patch: "0"}

var _ printer.TableFormattable = (*Version)(nil)
