package core

import (
	"bufio"
	"flag"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func ReadFromStdin() (input string, err error) {
	if flag.NArg() != 0 {
		input = flag.Arg(0)
		return input, nil
	}

	reader := bufio.NewReader(os.Stdin)

	input, err = reader.ReadString(';')
	if err != nil {
		log.Fatal("failed to read input")
	}

	input = strings.TrimSpace(input)
	log.Debug("File content:\n", input)

	return input, err
}
