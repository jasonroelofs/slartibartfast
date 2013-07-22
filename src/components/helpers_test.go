package components

// A helper object used in the components test suite.
// This file is named _test so it gets picked up by `go test` and ignored with
// `go build`

type TestHolder struct {
	Holding Component
	Removed []ComponentType
}

func (self *TestHolder) AddComponent(c Component) {
	self.Holding = c
}

func (self *TestHolder) GetComponent(t ComponentType) Component {
	return self.Holding
}

func (self *TestHolder) RemoveComponent(t ComponentType) {
	self.Removed = append(self.Removed, t)
}
