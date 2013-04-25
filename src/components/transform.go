package components

// Transform holds location and rotation data of the holding Entity
type Transform struct {
}

func (self Transform) Type() ComponentType {
	return TRANSFORM
}
