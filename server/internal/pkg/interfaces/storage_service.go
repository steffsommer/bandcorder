package interfaces

import (
	"server/internal/pkg/models"
	"time"
)

// StorageService saves raw audio data into an arbitrary audio file format and
// allows querying records created on certain days.
type StorageService interface {
	Save(fileName string, data []float32) error
	GetRecordings(date time.Time) ([]models.RecordingInfo, error)
	RenameRecording(oldFileName string, newFileName string, date time.Time) error
	DeleteRecording(fileName string, date time.Time) error
	RenameLastRecording(fileName string) error
}
