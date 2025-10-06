package services

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"server/internal/pkg/models"
	"server/internal/pkg/testutils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_FileSystemStorageService_SaveRecordingsAndGetFromTwoDays(t *testing.T) {
	tmpDir := t.TempDir()
	channelCount := 1
	sampleRate := 44100
	timeProvider := testutils.FakeTimeProvider{Time: time.Now()}
	service := NewFileSystemStorageService(tmpDir, channelCount, sampleRate, &timeProvider)
	file := "file.wav"
	data := make([]float32, sampleRate) // 1 second of audio
	err := service.Save(file, data)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	fakeTime := time.Date(2022, 1, 1, 23, 59, 59, 0, time.UTC)
	timeProvider.Time = fakeTime
	anotherFile := "another_file.wav"
	err = service.Save(anotherFile, data)
	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	todaysInfos, err := service.GetRecordings(time.Now())
	if err != nil {
		t.Fatalf("GetRecordings failed: %v", err)
	}
	expectedInfosToday := []models.RecordingInfo{{
		FileName:        file,
		DurationSeconds: 1,
	}}
	assert.Equal(t, expectedInfosToday, todaysInfos)

	otherDayInfos, err := service.GetRecordings(fakeTime.Truncate(time.Hour))
	if err != nil {
		t.Fatalf("GetRecordings failed: %v", err)
	}
	expectedOtherDayInfos := []models.RecordingInfo{{
		FileName:        anotherFile,
		DurationSeconds: 1,
	}}
	assert.Equal(t, expectedOtherDayInfos, otherDayInfos)
}

func Test_FileSystemStorageService_GetRecordings_Empty(t *testing.T) {
	tmpDir := t.TempDir()
	timeProvider := testutils.FakeTimeProvider{Time: time.Now()}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, &timeProvider)

	recordingInfos, err := service.GetRecordings(time.Now())

	if err != nil {
		t.Fatalf("GetRecordings failed: %v", err)
	}
	assert.Empty(t, recordingInfos)
}

func Test_CreatesWavFile_Correctly(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider)
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

func TestSave_CreatesDirectoryIfNotExists(t *testing.T) {
	tmpDir := t.TempDir()
	testTime := time.Date(2024, 3, 20, 14, 30, 0, 0, time.UTC)
	timeProvider := &testutils.FakeTimeProvider{Time: testTime}
	service := NewFileSystemStorageService(tmpDir, 1, 44100, timeProvider)
	expectedDir := filepath.Join(tmpDir, "2024-03-20")

	_, err := os.Stat(expectedDir)
	assert.True(t, os.IsNotExist(err))

	err = service.Save("test.wav", []float32{0.1})
	assert.NoError(t, err)

	stat, err := os.Stat(expectedDir)
	assert.NoError(t, err)
	assert.True(t, stat.IsDir())
}
