package interfaces

type AudioEffect int

const (
	MetronomeClick AudioEffect = iota
	RecordingStart
	RecordingStop
)

type PlaybackService interface {
	Play(effect AudioEffect)
}
