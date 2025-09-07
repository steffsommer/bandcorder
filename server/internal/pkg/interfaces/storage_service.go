package interfaces

import (
	"server/internal/pkg/models"
	"time"
)

type StorageService interface {
	Save(fileName string, data []float32) error
	GetRecordings(date time.Time) ([]models.RecordingInfo, error)
}
