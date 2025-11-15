package interfaces

type Metronome interface {
	Start() error
	UpdateBpm(bpm int) error
	Stop() error
}
