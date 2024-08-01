package markdown

import (
	_ "github.com/charmbracelet/glamour"
)

/*
Example:

	r, _ := glamour.NewTermRenderer(
	    // detect background color and pick either the default dark or light theme
	    glamour.WithAutoStyle(),
	    // wrap output at specific width (default is 80)
	    glamour.WithWordWrap(40),
	)

	out, _ := r.Render(`# Example`)

	fmt.Print(out)
*/
