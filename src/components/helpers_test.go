package components

type TestHolder struct {
	Holding Component
}

func (self *TestHolder) AddComponent(c Component) {
	self.Holding = c
}

func (self *TestHolder) GetComponent(t ComponentType) Component {
	return self.Holding
}

