package interfaces

import "time"

type RecordingMetaData struct {
	FileName string
	Started  time.Time
}

type Recorder interface {
	Start() (RecordingMetaData, error)
	Stop() error
	Abort() error
	GetMic() (string, error)
}
