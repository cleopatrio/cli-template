package cmd

type Version struct {
	Value string
}

func (T Version) Description() string {
	return T.Value
}

var version = Version{
	Value: "0.1.0",
}
