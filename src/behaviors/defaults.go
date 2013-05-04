package behaviors

type DataDefaults struct {
	LoadPath string
}

// Package-wide set of defaults, mainly so that things like data load
// paths are changeable in tests. Got the pattern here from go's log package
var defaults = New("data")

func New(loadPath string) *DataDefaults {
	return &DataDefaults{LoadPath: loadPath}
}
