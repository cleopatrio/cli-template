package printer

import (
	"fmt"
	"strings"

	"github.com/drewstinnett/gout/v2"
	"github.com/drewstinnett/gout/v2/config"
	"github.com/drewstinnett/gout/v2/formats/gotemplate"
	gJSON "github.com/drewstinnett/gout/v2/formats/json"
	gPlain "github.com/drewstinnett/gout/v2/formats/plain"
	gYAML "github.com/drewstinnett/gout/v2/formats/yaml"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/pflag"
)

type (
	OutputFormat struct {
		Options []string
		Default string
	}

	TableFormatter struct{}

	TableFormattable interface{ TableWriter() table.Writer }
)

var _ pflag.Value = (*OutputFormat)(nil)

func (ofe OutputFormat) String() string { return ofe.Default }

func (ofe *OutputFormat) Type() string { return "string" }

func (ofe *OutputFormat) Set(value string) error {
	isIncluded := func(opts []string, v string) bool {
		for _, opt := range opts {
			if v == opt {
				return true
			}
		}

		return false
	}

	if !isIncluded(ofe.Options, value) {
		return fmt.Errorf("%s is not a supported output format: %s", value, strings.Join(ofe.Options, ","))
	}

	ofe.Default = value
	return nil
}

func (f *TableFormatter) Format(data any) ([]byte, error) {
	tw, ok := data.(TableFormattable)
	if !ok {
		return []byte{}, nil
	}

	return []byte(tw.TableWriter().Render()), nil
}

var SetFormatter = func(printer *gout.Gout, format, template string) {
	switch format {
	case "table":
		printer.SetFormatter(&TableFormatter{})
	case "json":
		printer.SetFormatter(gJSON.Formatter{})
	case "yaml":
		printer.SetFormatter(gYAML.Formatter{})
	case "gotemplate":
		printer.SetFormatter(gotemplate.Formatter{
			Opts: config.FormatterOpts{"template": template},
		})
	default:
		printer.SetFormatter(gPlain.Formatter{})
	}
}
