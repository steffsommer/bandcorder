package models

import "time"

type EventId string

const (
	RecordingStateUpdate EventId = "RecordingState"
)

type EventLike interface {
	GetId() EventId
	GetData() any
}

type Event[DataT any] struct {
	EventId EventId
	Data    DataT `json:"data,omitempty"`
}

func (e Event[DataT]) GetId() EventId {
	return e.EventId
}

func (e Event[DataT]) GetData() any {
	return e.Data
}

type RecordingState string

const (
	IDLE    RecordingState = "IDLE"
	RUNNING RecordingState = "RUNNING"
)

type RecordingEventData struct {
	State    RecordingState `json:"state,omitempty"`
	FileName string         `json:"fileName,omitempty"`
	Started  time.Time      `json:"started,omitempty"`
}

func NewRecordingRunningEvent(
	fileName string,
	started time.Time,
) Event[RecordingEventData] {
	return Event[RecordingEventData]{
		EventId: RecordingStateUpdate,
		Data: RecordingEventData{
			State:    RUNNING,
			FileName: fileName,
			Started:  started,
		},
	}
}

func NewRecordingIdleEvent() Event[RecordingEventData] {
	return Event[RecordingEventData]{
		EventId: RecordingStateUpdate,
		Data: RecordingEventData{
			State: IDLE,
		},
	}
}
