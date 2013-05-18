package core

type Window interface {
	Open()
	IsOpen() bool
	SwapBuffers()
	Close()

	// AspectRatio returns the calculated aspect ratio of the current window
	AspectRatio() float32

	// TimeSinceLast returns the amount of time that has past in seconds
	// since the last time this method was called.
	TimeSinceLast() float32
}
