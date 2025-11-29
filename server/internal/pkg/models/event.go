package models

type EventId string

const (
	// recording events
	RecordingStartedEvent EventId = "RecordingStarted"
	RecordingStoppedEvent EventId = "RecordingStopped"
	RecordingAbortedEvent EventId = "RecordingAborted"
	RecordingDeletedEvent EventId = "RecordingDeleted"
	RecordingRenamedEvent EventId = "FileRenamed"
	// continously broadcasted recording events
	RecordingIdleEvent    EventId = "RecordingIdle"
	RecordingRunningEvent EventId = "RecordingRunning"
	// metronome events
	MetronomeBeatEvent        EventId = "MetronomeBeat"
	MetronomeStateChangeEvent EventId = "MetronomeStateChange"
	// misc
	LiveAudioDataEvent   EventId = "LiveAudioData"
	SettingsUpdatedEvent EventId = "SettingsUpdated"
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
