package helpers

import (
	"io"
	"os"
	"strings"
)

func ReadStdin() (input string, err error) {
	stdin, err := io.ReadAll(os.Stdin)
	if err != nil {
		return "", err
	}

	content := strings.TrimSpace(string(stdin))

	return content, err
}
