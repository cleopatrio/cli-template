package core

import (
	"sync"
	"time"

	"clitemplate/cmd/cli/core/formatters"
	"clitemplate/pkg/dbsql"

	"github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats/gotemplate"
	gJSON "github.com/drewstinnett/gout/v2/formats/json"
	gYAML "github.com/drewstinnett/gout/v2/formats/yaml"
	"github.com/oleoneto/go-toolkit/files"
	"github.com/spf13/cobra"
)

type CommandFlags struct {
	VerboseLogging  bool
	OutputTemplate  string
	OutputFormat    *FlagEnum
	Engine          *FlagEnum
	Directory       string
	DatabaseName    string
	DatabaseURL     *string
	TimeExecutions  bool
	ServerAddr      string
	DevelopmentMode bool
	HomeDirectory   string
	CLIConfig       string
	ConfigDir       files.File
	File            FileFlag
	Stdin           string
}

type CommandState struct {
	Writer             *gout.Gout
	Flags              CommandFlags
	ExecutionStartTime time.Time
	ExecutionExitLog   []any
	Database           dbsql.SqlBackend
}

var cliDir = ".clitempla"

var defaultFlags = CommandFlags{
	DatabaseName:    "clitemplate.sqlite",
	DevelopmentMode: false,
	OutputFormat: &FlagEnum{
		Allowed: []string{"plain", "json", "yaml", "table", "gotemplate", "silent"},
		Default: "plain",
	},
	Engine: &FlagEnum{
		Allowed: []string{"postgresql", "sqlite3"},
		Default: "sqlite3",
	},
	ServerAddr: "0.0.0.0:40400",
	ConfigDir: files.NewDirectory(
		cliDir,
		files.File{Name: "config.yaml"},
		files.File{Name: "cache", IsDirectory: true},
		files.File{Name: "data", IsDirectory: true},
	),
}

var once sync.Once
var state *CommandState

func NewCommandState() *CommandState {
	once.Do(func() {
		state = &CommandState{
			Writer: gout.New(),
			Flags:  defaultFlags,
		}
	})

	return state
}

func (c *CommandState) SetFormatter(cmd *cobra.Command, args []string) {
	switch cmd.Flag("output").Value.String() {
	case "table":
		c.Writer.SetFormatter(&formatters.TableFormatter{})
	case "json":
		c.Writer.SetFormatter(gJSON.Formatter{})
	case "yaml":
		c.Writer.SetFormatter(gYAML.Formatter{})
	case "gotemplate":
		c.Writer.SetFormatter(gotemplate.Formatter{
			Opts: config.FormatterOpts{"template": c.Flags.OutputTemplate},
		})
	case "silent":
		c.Writer.SetFormatter(formatters.SilentFormatter{})
	case "plain":
		c.Writer.SetFormatter(formatters.PlainFormatter{})
	}
}
