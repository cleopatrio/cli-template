package core

import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

type FlagEnum struct {
	Allowed []string
	Default string
}

func (ofe FlagEnum) String() string { return ofe.Default }

func (ofe *FlagEnum) Type() string { return "string" }

func (ofe *FlagEnum) Set(value string) error {
	isIncluded := func(opts []string, v string) bool {
		for _, opt := range opts {
			if v == opt {
				return true
			}
		}

		return false
	}

	if !isIncluded(ofe.Allowed, value) {
		return fmt.Errorf("%s is not a supported output format: %s", value, strings.Join(ofe.Allowed, ","))
	}

	ofe.Default = value
	return nil
}

var _ pflag.Value = (*FlagEnum)(nil)
