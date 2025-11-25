package models

type EventId string

const (
	RecordingIdleEvent        EventId = "RecordingIdle"
	RecordingRunningEvent     EventId = "RecordingRunning"
	LiveAudioDataEvent        EventId = "LiveAudioData"
	FileRenamedEvent          EventId = "FileRenamed"
	MetronomeBeatEvent        EventId = "MetronomeBeat"
	MetronomeStateChangeEvent EventId = "MetronomeStateChange"
	SettingsUpdatedEvent      EventId = "SettingsUpdated"
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
