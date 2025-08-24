package interfaces

type Recorder interface {
	Start() error
	Stop() error
	Abort() error
}
