package services

import (
	"os"
	"path/filepath"

	"github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
)

const fileName = "config.yml"
const configFolder = "bandcorder"

type Settings struct {
	RecordingsDirectory string
}

type SettingsService struct {
}

func NewSettingsService() *SettingsService {
	return &SettingsService{}
}

func (s *SettingsService) Load() (Settings, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		logrus.WithError(err).Error("Failed to determine user config directory")
		return Settings{}, err
	}

	defaults, err := s.getDefaultSettings()
	if err != nil {
		logrus.WithError(err).Error("Failed to determine user home directory")
	}

	configFilePath := filepath.Join(configDir, configFolder, fileName)
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(configFilePath), 0755)
		_, err := os.Create(configFilePath)
		if err != nil {
			logrus.WithError(err).Errorf("Failed to create empty config file %s", configFilePath)
			return Settings{}, err
		}
	}

	yamlContent, err := os.ReadFile(configFilePath)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to read config file %s", configFilePath)
		return Settings{}, err
	}

	var userSettings Settings
	err = yaml.Unmarshal(yamlContent, &userSettings)
	if err != nil {
		logrus.WithError(err).Error("Config file has an invalid format")
		return Settings{}, err
	}

	settings := s.mergeSettings(userSettings, defaults)
	return settings, nil
}

func (s *SettingsService) getDefaultSettings() (Settings, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Failed to determine user home directory")
		return Settings{}, err
	}
	return Settings{
		RecordingsDirectory: filepath.Join(homeDir, "Documents", "Recordings"),
	}, nil
}

func (s *SettingsService) mergeSettings(defaults, userSettings Settings) Settings {
	if userSettings.RecordingsDirectory != "" {
		defaults.RecordingsDirectory = userSettings.RecordingsDirectory
	}
	return defaults
}
