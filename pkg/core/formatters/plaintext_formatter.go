package formatters

import "fmt"

type Stringable interface{ String() string }

type PlainFormatter struct{}

func (w PlainFormatter) Format(v interface{}) ([]byte, error) {
	s, ok := v.(Stringable)
	if !ok {
		return []byte(fmt.Sprintf("%+v", v)), nil
	}

	return []byte(s.String()), nil
}
