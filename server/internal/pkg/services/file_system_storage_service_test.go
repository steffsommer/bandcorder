package services

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"server/internal/pkg/models"
	"server/internal/pkg/testutils"
	"server/internal/pkg/testutils/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_SaveRecordingsAndGetFromTwoDays(t *testing.T) {
	tmpDir := t.TempDir()
	channelCount := 1
	sampleRate := 44100
	timeProvider := testutils.FakeTimeProvider{Time: time.Now()}
	service := NewFileSystemStorageService(tmpDir, channelCount, sampleRate, &timeProvider, nil)
	file := "file.wav"
	data := make([]float32, sampleRate)
	err := service.Save(file, data)
	assert.NoError(t, err)

	fakeTime := time.Date(2022, 1, 1, 23, 59, 59, 0, time.UTC)
	timeProvider.Time = fakeTime
	anotherFile := "another_file.wav"
	err = service.Save(anotherFile, data)
	assert.NoError(t, err)

	todaysInfos, err := service.GetRecordings(time.Now())
	assert.NoError(t, err)
	assert.Len(t, todaysInfos, 1)
	assert.Equal(t, file, todaysInfos[0].FileName)
	assert.Equal(t, uint32(1), todaysInfos[0].DurationSeconds)

	otherDayInfos, err := service.GetRecordings(fakeTime.Truncate(time.Hour))
	assert.NoError(t, err)
	assert.Len(t, otherDayInfos, 1)
	assert.Equal(t, anotherFile, otherDayInfos[0].FileName)
	assert.Equal(t, uint32(1), otherDayInfos[0].DurationSeconds)
}

func Test_GetRecordings_Empty(t *testing.T) {
	tmpDir := t.TempDir()
	timeProvider := testutils.FakeTimeProvider{Time: time.Now()}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, &timeProvider, nil)

	recordingInfos, err := service.GetRecordings(time.Now())
	assert.NoError(t, err)
	assert.Empty(t, recordingInfos)
}

func Test_GetRecordings_OrderedByModTimeDesc(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)

	// Create files with different mod times
	err := service.Save("first.wav", []float32{0.1})
	assert.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	err = service.Save("second.wav", []float32{0.2})
	assert.NoError(t, err)

	time.Sleep(10 * time.Millisecond)

	err = service.Save("third.wav", []float32{0.3})
	assert.NoError(t, err)

	recordings, err := service.GetRecordings(testTime)
	assert.NoError(t, err)
	assert.Len(t, recordings, 3)
	assert.Equal(t, "third.wav", recordings[0].FileName)
	assert.Equal(t, "second.wav", recordings[1].FileName)
	assert.Equal(t, "first.wav", recordings[2].FileName)
}

func Test_CreatesWavFile_Correctly(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)
	testData := []float32{0.1, 0.2, 0.3, 0.4}

	err := service.Save("test.wav", testData)
	assert.NoError(t, err)

	expectedPath := filepath.Join(tmpDir, "2024-01-15", "test.wav")
	file, err := os.Open(expectedPath)
	assert.NoError(t, err)
	defer file.Close()

	var header WAVHeader
	binary.Read(file, binary.LittleEndian, &header)
	assert.Equal(t, "RIFF", string(header.ChunkID[:]))
	assert.Equal(t, uint32(44100), header.SampleRate)

	audioData := make([]float32, len(testData))
	binary.Read(file, binary.LittleEndian, &audioData)
	assert.Equal(t, testData, audioData)
}

func Test_Save_CreatesDirectoryIfNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 3, 20, 14, 30, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)
	expectedDir := filepath.Join(tmpDir, "2024-03-20")

	_, err := os.Stat(expectedDir)
	assert.True(t, os.IsNotExist(err))

	err = service.Save("test.wav", []float32{0.1})
	assert.NoError(t, err)

	stat, err := os.Stat(expectedDir)
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}

func Test_RenameRecording_Success(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)

	testData := []float32{0.1, 0.2}
	err := service.Save("old_name.wav", testData)
	assert.NoError(t, err)

	err = service.RenameRecording("old_name.wav", "new_name.wav", testTime)
	assert.NoError(t, err)

	dateDir := filepath.Join(tmpDir, "2024-01-15")
	oldPath := filepath.Join(dateDir, "old_name.wav")
	newPath := filepath.Join(dateDir, "new_name.wav")

	_, err = os.Stat(oldPath)
	assert.True(t, os.IsNotExist(err))

	_, err = os.Stat(newPath)
	assert.NoError(t, err)
}

func Test_RenameRecording_DateDirNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	timeProvider := &testutils.FakeTimeProvider{Time: time.Now()}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)

	nonExistentDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	err := service.RenameRecording("old.wav", "new.wav", nonExistentDate)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No recordings exist for given date")
}

func Test_RenameRecording_FileNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)

	testData := []float32{0.1, 0.2}
	err := service.Save("existing.wav", testData)
	assert.NoError(t, err)

	err = service.RenameRecording("nonexistent.wav", "new.wav", testTime)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func Test_DeleteRecording_Success(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)

	err := service.Save("test.wav", []float32{0.1, 0.2})
	assert.NoError(t, err)

	err = service.DeleteRecording("test.wav", testTime)
	assert.NoError(t, err)

	filePath := filepath.Join(tmpDir, "2024-01-15", "test.wav")
	_, err = os.Stat(filePath)
	assert.True(t, os.IsNotExist(err))
}

func Test_DeleteRecording_DateDirNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	timeProvider := &testutils.FakeTimeProvider{Time: time.Now()}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)

	nonExistentDate := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	err := service.DeleteRecording("test.wav", nonExistentDate)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "No recordings exist for given date")
}

func Test_DeleteRecording_FileNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider, nil)

	err := service.Save("existing.wav", []float32{0.1})
	assert.NoError(t, err)

	err = service.DeleteRecording("nonexistent.wav", testTime)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "does not exist")
}

func createTestDirStructure(t *testing.T, baseDir string, structure map[string][]string) {
	for dir, files := range structure {
		dirPath := filepath.Join(baseDir, dir)
		err := os.MkdirAll(dirPath, 0755)
		assert.NoError(t, err)
		for _, file := range files {
			filePath := filepath.Join(dirPath, file)
			err := os.WriteFile(filePath, []byte("test"), 0644)
			assert.NoError(t, err)
			time.Sleep(5 * time.Millisecond)
		}
	}
}

func Test_RenameLastRecording(t *testing.T) {
	tmpDir := t.TempDir()

	structure := map[string][]string{
		"2024-01-01": {"old1.wav"},
		"2024-01-03": {"old2.wav", "old3.wav"},
		"2024-01-02": {},
	}
	createTestDirStructure(t, tmpDir, structure)

	timeProvider := testutils.FakeTimeProvider{Time: time.Now()}
	eventBus := mocks.NewMockEventBus(t)
	service := NewFileSystemStorageService(tmpDir, 1, 44100, &timeProvider, eventBus)
	eventBus.EXPECT().Dispatch(models.NewRecordingRenamedEvent())

	err := service.RenameLastRecording("renamed.wav")
	assert.NoError(t, err)

	files, err := os.ReadDir(filepath.Join(tmpDir, "2024-01-03"))
	assert.NoError(t, err)

	var fileNames []string
	for _, f := range files {
		fileNames = append(fileNames, f.Name())
	}
	assert.Contains(t, fileNames, "renamed.wav")
	assert.NotContains(t, fileNames, "old3.wav")
	assert.Contains(t, fileNames, "old2.wav")
}

func Test_RenameLastRecording_NoFiles(t *testing.T) {
	tmpDir := t.TempDir()
	structure := map[string][]string{
		"2024-01-01": {},
		"2024-01-02": {},
	}
	createTestDirStructure(t, tmpDir, structure)
	service := &FileSystemStorageService{baseDir: tmpDir}
	err := service.RenameLastRecording("shouldnotmatter.wav")
	assert.Error(t, err)
}

func Test_UpdatesRecordingDirectory_OnSettingsChange(t *testing.T) {
	tempDir1 := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	eventBus := &ManualMockEventBus{}
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tempDir1, 1, 44100, timeProvider, eventBus)
	assert.Equal(t, tempDir1, service.baseDir)
}

func Test_FailsSave_WhenRecordingDirectoryDoesNotExist(t *testing.T) {
	tempDir := filepath.Join(os.TempDir(), "non_existing_dir")
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	eventBus := &ManualMockEventBus{}
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tempDir, 1, 44100, timeProvider, eventBus)
	err := service.Save("first.wav", []float32{0.1})
	assert.Error(t, err)
	assert.ErrorContains(t, err, "not exist")
}

type ManualMockEventBus struct {
	callbacks []func(any)
}

func (m *ManualMockEventBus) Dispatch(event models.EventLike) {
	for _, cb := range m.callbacks {
		cb(event)
	}
}

func (m *ManualMockEventBus) OnEvent(eventId models.EventId, cb func(any)) {
	m.callbacks = append(m.callbacks, cb)
}
