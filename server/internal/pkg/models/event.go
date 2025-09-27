package models

import "time"

type EventId string

const (
	RecordingIdleEvent    EventId = "RecordingIdle"
	RecordingRunningEvent EventId = "RecordingRunning"
)

type EventLike interface {
	GetId() EventId
	GetData() any
}

type Event[DataT any] struct {
	EventId EventId `json:"eventId"`
	Data    DataT   `json:"data,omitempty"`
}

func (e Event[DataT]) GetId() EventId {
	return e.EventId
}

func (e Event[DataT]) GetData() any {
	return e.Data
}

type RecordingEventData struct {
	FileName string    `json:"fileName,omitempty"`
	Started  time.Time `json:"started,omitempty"`
}

func NewRecordingRunningEvent(
	fileName string,
	started time.Time,
) Event[RecordingEventData] {
	return Event[RecordingEventData]{
		EventId: RecordingRunningEvent,
		Data: RecordingEventData{
			FileName: fileName,
			Started:  started,
		},
	}
}

func NewRecordingIdleEvent() Event[RecordingEventData] {
	return Event[RecordingEventData]{
		EventId: RecordingIdleEvent,
	}
}
