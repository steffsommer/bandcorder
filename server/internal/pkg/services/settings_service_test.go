package services

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/goccy/go-yaml"
	"github.com/stretchr/testify/assert"
)

func Test_CreateDefaultSettingsFile_OnLoad(t *testing.T) {
	tempDir := createTempDir(t)
	tempFilePath := filepath.Join(tempDir, "config.yaml")
	service := NewSettingsService(tempFilePath)
	_, err := service.Load()
	assert.NoError(t, err)
	content, err := os.ReadFile(tempFilePath)
	assert.Nil(t, err)
	var settings Settings
	err = yaml.Unmarshal(content, &settings)
	assert.Nil(t, err)
	assert.NotEmpty(t, settings.RecordingsDirectory)
}

func Test_Succeed_Loading(t *testing.T) {
	tempFile := createTempFile(t)
	fmt.Println(tempFile)
	service := NewSettingsService(tempFile)
	settings := Settings{
		RecordingsDirectory: "/tmp/recordings",
	}
	settingsYaml, _ := yaml.Marshal(&settings)
	os.WriteFile(tempFile, settingsYaml, 0755)
	loadedSettings, err := service.Load()
	assert.NoError(t, err)
	assert.Equal(t, settings.RecordingsDirectory, loadedSettings.RecordingsDirectory)
	os.ReadFile(tempFile)
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
