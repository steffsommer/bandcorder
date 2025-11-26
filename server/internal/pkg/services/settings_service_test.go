package services

import (
	"os"
	"path/filepath"
	"server/internal/pkg/models"
	"server/internal/pkg/testutils/mocks"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
)

func Test_CreateDefaultSettingsFile_OnLoad(t *testing.T) {
	tempDir := createTempDir(t)
	tempFilePath := filepath.Join(tempDir, "config.yaml")
	dispatcher := mocks.NewMockEventDispatcher(t)
	service := NewSettingsService(tempFilePath, dispatcher)
	_, err := service.Load()
	assert.NoError(t, err)
	content, err := os.ReadFile(tempFilePath)
	assert.Nil(t, err)
	var settings models.Settings
	err = yaml.Unmarshal(content, &settings)
	assert.Nil(t, err)
	assert.NotEmpty(t, settings.RecordingsDirectory)
}

func Test_Succeed_Loading(t *testing.T) {
	tempFile := createTempFile(t)
	dispatcher := mocks.NewMockEventDispatcher(t)
	service := NewSettingsService(tempFile, dispatcher)
	settings := models.Settings{
		RecordingsDirectory: "/tmp/recordings",
	}
	settingsYaml, _ := yaml.Marshal(&settings)
	os.WriteFile(tempFile, settingsYaml, 0755)
	loadedSettings, err := service.Load()
	assert.NoError(t, err)
	assert.Equal(t, settings.RecordingsDirectory, loadedSettings.RecordingsDirectory)
	os.ReadFile(tempFile)
}

func Test_CallsOnUpdateCallback(t *testing.T) {
	tempDir := createTempDir(t)
	tempFilePath := filepath.Join(tempDir, "config.yaml")
	settings := models.Settings{
		RecordingsDirectory: "/some/dir",
	}
	dispatcher := mocks.NewMockEventDispatcher(t)
	service := NewSettingsService(tempFilePath, dispatcher)
	dispatcher.EXPECT().Dispatch(models.NewSettingsUpdatedEvent(settings))
	err := service.Save(settings)
	assert.NoError(t, err)
}

func Test_SaveSettings(t *testing.T) {
	tempDir := createTempDir(t)
	tempFilePath := filepath.Join(tempDir, "config.yaml")
	settings := models.Settings{
		RecordingsDirectory: "/some/dir",
	}
	dispatcher := mocks.NewMockEventDispatcher(t)
	service := NewSettingsService(tempFilePath, dispatcher)
	dispatcher.EXPECT().Dispatch(models.NewSettingsUpdatedEvent(settings))
	err := service.Save(settings)
	assert.NoError(t, err)
	loadedSettings, err := service.Load()
	assert.NoError(t, err)
	assert.Equal(t, settings, loadedSettings)
}

func createTempFile(t *testing.T) string {
	file, err := os.CreateTemp("", "config.yaml")
	if err != nil {
		panic("Failed to create temp file: " + err.Error())
	}
	path, err := filepath.Abs(file.Name())
	if err != nil {
		panic("Failed to obtain temp file absolute path: " + err.Error())
	}
	t.Cleanup(func() { os.Remove(path) })
	return path
}

func createTempDir(t *testing.T) string {
	tempDir, err := os.MkdirTemp("", "bandcorder-")
	assert.NoError(t, err)
	t.Cleanup(func() { os.Remove(tempDir) })
	return tempDir
}

func noOpSettingsConsumer(s models.Settings) {}
