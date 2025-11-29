package models

import "time"

type RecordingRunningEventData struct {
	FileName       string `json:"fileName"`
	SecondsRunning uint32 `json:"secondsRunning"`
}

type RecordingStartedEventData struct {
	FileName string
	Started  time.Time
}

func NewRecordingStartedEvent(
	fileName string,
	started time.Time,
) Event[RecordingStartedEventData] {
	return Event[RecordingStartedEventData]{
		EventId: RecordingStartedEvent,
		Data: RecordingStartedEventData{
			FileName: fileName,
			Started:  started,
		},
	}
}

func NewRecordingStoppedEvent() Event[any] {
	return Event[any]{
		EventId: RecordingStoppedEvent,
	}
}

func NewRecordingAbortedEvent() Event[any] {
	return Event[any]{
		EventId: RecordingAbortedEvent,
	}
}

func NewRecordingDeletedEvent() Event[any] {
	return Event[any]{
		EventId: RecordingDeletedEvent,
	}
}

func NewRecordingRenamedEvent() Event[any] {
	return Event[any]{
		EventId: RecordingRenamedEvent,
	}
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
