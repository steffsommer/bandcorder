package models

import "time"

type EventId string

const (
	RecordingIdleEvent    EventId = "RecordingIdle"
	RecordingRunningEvent EventId = "RecordingRunning"
	LiveAudioDataEvent    EventId = "LiveAudioData"
	FileRenamedEvent      EventId = "FileRenamed"
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

type RunningEventData struct {
	FileName       string `json:"fileName"`
	SecondsRunning uint32 `json:"secondsRunning"`
}

func NewRecordingRunningEvent(
	fileName string,
	started time.Time,
) Event[RunningEventData] {
	duration := time.Since(started)
	return Event[RunningEventData]{
		EventId: RecordingRunningEvent,
		Data: RunningEventData{
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

type LiveAudioEventData struct {
	LoudnessPercentage uint8 `json:"loudnessPercentage"`
	FrequencyBars      []int `json:"frequencyBars"`
}

func NewLiveAudioDataEvent(
	loudnessPercentage uint8,
	frequencyBars []int,
) Event[LiveAudioEventData] {
	return Event[LiveAudioEventData]{
		EventId: LiveAudioDataEvent,
		Data: LiveAudioEventData{
			LoudnessPercentage: loudnessPercentage,
			FrequencyBars:      frequencyBars,
		},
	}
}

func NewFileRenamedEvent() Event[any] {
	return Event[any]{
		EventId: FileRenamedEvent,
	}
}
