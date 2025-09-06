package interfaces

import "time"

type StartedResponse struct {
	FileName string
	Started  time.Time
}

type Recorder interface {
	Start() (StartedResponse, error)
	Stop() error
	Abort() error
}
