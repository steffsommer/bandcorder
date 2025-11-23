package facades

import (
	"errors"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"server/internal/pkg/testutils/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	mockPlayer  *mocks.MockPlaybackService
	mockStorage *mocks.MockStorageService
)

func setupStorageFacade(t *testing.T) *FileSystemStorageFacade {
	mockPlayer = mocks.NewMockPlaybackService(t)
	mockStorage = mocks.NewMockStorageService(t)
	return NewFileSystemStorageFacade(mockPlayer, mockStorage)
}

func Test_SavesFile_Success(t *testing.T) {
	facade := setupStorageFacade(t)
	mockStorage.EXPECT().Save("test.wav", []float32{1.0, 2.0}).Return(nil)

	err := facade.Save("test.wav", []float32{1.0, 2.0})

	assert.NoError(t, err)
}

func Test_GetRecordings(t *testing.T) {
	facade := setupStorageFacade(t)
	date := time.Now()
	expected := []models.RecordingInfo{{FileName: "test.wav"}}
	mockStorage.EXPECT().GetRecordings(date).Return(expected, nil)

	recordings, err := facade.GetRecordings(date)

	assert.NoError(t, err)
	assert.Equal(t, expected, recordings)
}

func Test_RenameRecording_Success(t *testing.T) {
	facade := setupStorageFacade(t)
	date := time.Now()
	mockStorage.EXPECT().RenameRecording("old.wav", "new.wav", date).Return(nil)
	mockPlayer.EXPECT().Play(interfaces.Success).Return()

	err := facade.RenameRecording("old.wav", "new.wav", date)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func Test_RenameRecording_Error(t *testing.T) {
	facade := setupStorageFacade(t)
	date := time.Now()
	mockStorage.EXPECT().RenameRecording("old.wav", "new.wav", date).Return(errors.New("rename failed"))

	err := facade.RenameRecording("old.wav", "new.wav", date)

	assert.Error(t, err)
}

func Test_DeleteRecording_Success(t *testing.T) {
	facade := setupStorageFacade(t)
	date := time.Now()
	mockStorage.EXPECT().DeleteRecording("test.wav", date).Return(nil)
	mockPlayer.EXPECT().Play(interfaces.Delete).Return()

	err := facade.DeleteRecording("test.wav", date)

	assert.NoError(t, err)
}

func Test_DeleteRecording_Error(t *testing.T) {
	facade := setupStorageFacade(t)
	date := time.Now()
	mockStorage.EXPECT().DeleteRecording("test.wav", date).Return(errors.New("delete failed"))

	err := facade.DeleteRecording("test.wav", date)

	assert.Error(t, err)
}

func Test_RenameLastRecording_Success(t *testing.T) {
	facade := setupStorageFacade(t)
	mockStorage.EXPECT().RenameLastRecording("new.wav").Return(nil)
	mockPlayer.EXPECT().Play(interfaces.Success).Return()

	err := facade.RenameLastRecording("new.wav")

	assert.NoError(t, err)
}

func Test_RenameLastRecording_Error(t *testing.T) {
	facade := setupStorageFacade(t)
	mockStorage.EXPECT().RenameLastRecording("new.wav").Return(errors.New("rename failed"))

	err := facade.RenameLastRecording("new.wav")

	assert.Error(t, err)
}
