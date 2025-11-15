package facades

import (
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"time"
)

type FileSystemStorageFacade struct {
	player         interfaces.PlaybackService
	storageService interfaces.StorageService
}

func NewFileSystemStorageFacade(
	player interfaces.PlaybackService,
	storageService interfaces.StorageService,
) *FileSystemStorageFacade {
	return &FileSystemStorageFacade{
		player:         player,
		storageService: storageService,
	}
}

func (f *FileSystemStorageFacade) Save(fileName string, data []float32) error {
	return f.storageService.Save(fileName, data)
}

func (f *FileSystemStorageFacade) GetRecordings(date time.Time) ([]models.RecordingInfo, error) {
	return f.storageService.GetRecordings(date)
}

func (f *FileSystemStorageFacade) RenameRecording(oldFileName string, newFileName string, date time.Time) error {
	err := f.storageService.RenameRecording(oldFileName, newFileName, date)
	if err == nil {
		f.player.Play(interfaces.Success)
	}
	return err
}

func (f *FileSystemStorageFacade) DeleteRecording(fileName string, date time.Time) error {
	err := f.storageService.DeleteRecording(fileName, date)
	if err == nil {
		f.player.Play(interfaces.Delete)
	}
	return err
}

func (f *FileSystemStorageFacade) RenameLastRecording(fileName string) error {
	err := f.storageService.RenameLastRecording(fileName)
	if err == nil {
		f.player.Play(interfaces.Success)
	}
	return err
}
