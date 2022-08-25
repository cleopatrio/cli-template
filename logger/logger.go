package logger

import (
	"fmt"
	"io"
	"strings"

	"github.com/drewstinnett/go-output-format/formatter"
)

type Logger struct {
	config         formatter.Config
	cachedMessages []Formattable
}

type Formattable interface {
	Description() string
}

type ApplicationMessage struct {
	Message string `json:"message" yaml:"message"`
}

type ApplicationMessages []ApplicationMessage

type ApplicationError struct {
	Error string `json:"error" yaml:"error"`
}

func (T Logger) Log(data Formattable, ioWriter io.Writer) {
	output, err := formatter.OutputData(data, &T.config)

	if err != nil {
		fmt.Fprintln(ioWriter, err)
		return
	}

	switch T.config.Format {
	case "plain":
		fmt.Fprintln(ioWriter, data.Description())
	default:
		fmt.Fprintln(ioWriter, string(output))
	}
}

func (T *Logger) CacheMessage(message Formattable) {
	T.cachedMessages = append(T.cachedMessages, message)
}

func (T *Logger) ReleaseCachedMessages(ioWriter io.Writer) {
	switch T.config.Format {
	case "plain":
		messages := []string{}

		for _, m := range T.cachedMessages {
			messages = append(messages, m.Description())
		}

		fmt.Fprintln(ioWriter, strings.Join(messages, "\n"))
	default:
		output, err := formatter.OutputData(T.cachedMessages, &T.config)

		if err != nil {
			fmt.Fprintln(ioWriter, err)
			return
		}

		fmt.Fprintln(ioWriter, string(output))
	}
}

func (T ApplicationError) Description() string {
	return T.Error
}

func (T ApplicationMessage) Description() string {
	return T.Message
}

// MARK: Default Loggers

func Default() Logger {
	return Logger{
		config: formatter.Config{
			Template: "",
			Format:   "plain",
		},
	}
}

func Custom(format, template string) Logger {
	return Logger{
		config: formatter.Config{
			Template: template,
			Format:   format,
		},
	}
}
