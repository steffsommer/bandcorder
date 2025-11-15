package models

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
