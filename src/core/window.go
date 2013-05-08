package core

type Window interface {
	Open()
	IsOpen() bool
	SwapBuffers()
	Close()

	// TimeSinceLast returns the amount of time that has past in seconds
	// since the last time this method was called.
	TimeSinceLast() float64
}
