package eventbus

import "time"

type recordingRunningEvent struct {
	FileName string
	Started  time.Time
}
