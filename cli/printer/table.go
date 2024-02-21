package printer

import (
	"io"

	"github.com/jedib0t/go-pretty/v6/table"
)

type (
	TableRowOptions struct {
		Simplified bool
	}

	TableOptions struct {
		ColumnConfig *[]table.ColumnConfig
		Title        *string
		Header       table.Row
		Footer       *table.Row
		Style        *table.Style
		OutputMirror io.Writer
	}
)

func Initialize(opts TableOptions) table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(opts.OutputMirror)

	if opts.Style != nil {
		t.SetStyle(*opts.Style)
	} else {
		t.SetStyle(table.StyleLight)
	}

	if opts.ColumnConfig != nil {
		t.SetColumnConfigs(*opts.ColumnConfig)
	}

	if opts.Title != nil {
		t.SetTitle(*opts.Title)
	}

	t.AppendHeader(opts.Header)

	if opts.Footer != nil {
		t.AppendFooter(*opts.Footer)
	}

	return t
}
