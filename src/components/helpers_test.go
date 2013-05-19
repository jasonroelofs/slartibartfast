package components

// A helper object used in the components test suite.
// This file is named _test so it gets picked up by `go test` and ignored with
// `go build`

type TestHolder struct {
	Holding Component
}

func (self *TestHolder) AddComponent(c Component) {
	self.Holding = c
}

func (self *TestHolder) GetComponent(t ComponentType) Component {
	return self.Holding
}

