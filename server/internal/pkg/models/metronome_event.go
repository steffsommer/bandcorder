package models

type MetronomeBeatEventData struct {
	BeatCount int `json:"beatCount"`
}

func NewMetronomeBeatEvent(beatCount int) Event[MetronomeBeatEventData] {
	return Event[MetronomeBeatEventData]{
		EventId: MetronomeRunningEvent,
		Data: MetronomeBeatEventData{
			BeatCount: beatCount,
		},
	}
}

func NewMetronomeIdleEvent() Event[any] {
	return Event[any]{
		EventId: MetronomeIdleEvent,
	}
}
