package interfaces

type Metronome interface {
	Start()
	UpdateBpm(bpm int)
	Stop()
}
