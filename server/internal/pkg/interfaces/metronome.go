package interfaces

import "server/internal/pkg/models"

type Metronome interface {
	Start() error
	UpdateBpm(bpm int) error
	Stop() error
	GetState() models.MetronomeStateEventData
}
