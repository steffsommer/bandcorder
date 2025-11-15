package interfaces

type AudioEffect int

const (
	MetronomeClick AudioEffect = iota
	SwitchOn
	SwitchOff
	Delete
	Success
)

type PlaybackService interface {
	Play(effect AudioEffect)
}
