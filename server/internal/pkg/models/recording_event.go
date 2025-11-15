package models

import "time"

type RecordingRunningEventData struct {
	FileName       string `json:"fileName"`
	SecondsRunning uint32 `json:"secondsRunning"`
}

func NewRecordingRunningEvent(
	fileName string,
	started time.Time,
) Event[RecordingRunningEventData] {
	duration := time.Since(started)
	return Event[RecordingRunningEventData]{
		EventId: RecordingRunningEvent,
		Data: RecordingRunningEventData{
			FileName:       fileName,
			SecondsRunning: uint32(duration.Seconds()),
		},
	}
}

func NewRecordingIdleEvent() Event[any] {
	return Event[any]{
		EventId: RecordingIdleEvent,
	}
}
