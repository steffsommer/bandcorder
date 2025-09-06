package eventbus

import "time"

type RecordingState string

const (
	IDLE    RecordingState = "IDLE"
	RUNNING RecordingState = "RUNNING"
)

type recordingRunningEvent struct {
	State    RecordingState
	FileName string
	Started  time.Time
}
