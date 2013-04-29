package core

type Window interface {
	Open()
	IsOpen() bool
	SwapBuffers()
	Close()
}
