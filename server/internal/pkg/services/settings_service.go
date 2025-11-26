package services

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"server/internal/pkg/utils"

	"github.com/goccy/go-yaml"
	"github.com/sirupsen/logrus"
)

const fileName = "config.yml"
const configFolder = "bandcorder"

// SettingsService reads and writes the application settings to a YAML file
type SettingsService struct {
	filePath   string
	dispatcher interfaces.EventDispatcher
}

// NewSettingsService creates a new SettingsService
func NewSettingsService(filePath string, dispatcher interfaces.EventDispatcher) *SettingsService {
	return &SettingsService{
		filePath:   filePath,
		dispatcher: dispatcher,
	}
}

// Load loads the application settings from disk. If the settings file does not exist,
// a new one with the default settings will be created.
func (s *SettingsService) Load() (models.Settings, error) {
	defaults, err := s.getDefaults()
	if err != nil {
		return models.Settings{}, err
	}
	s.createFileIfMissing(defaults)
	rawSettings, err := os.ReadFile(s.filePath)
	if err != nil {
		logrus.WithError(err).Errorf("Failed to read config file %s", s.filePath)
		return models.Settings{}, err
	}
	var userSettings models.Settings
	err = yaml.Unmarshal(rawSettings, &userSettings)
	if err != nil {
		logrus.WithError(err).Error("Config file has an invalid format")
		return models.Settings{}, err
	}
	settings := s.merge(defaults, userSettings)
	return settings, nil
}

func (s *SettingsService) getDefaults() (models.Settings, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		logrus.WithError(err).Error("Failed to determine user home directory")
		return models.Settings{}, err
	}
	return models.Settings{
		RecordingsDirectory: filepath.Join(homeDir, "Documents", "Recordings"),
	}, nil
}

func (s *SettingsService) merge(defaults, userSettings models.Settings) models.Settings {
	if userSettings.RecordingsDirectory != "" {
		defaults.RecordingsDirectory = userSettings.RecordingsDirectory
	}
	return defaults
}

// Saves saves the settings to disk
func (s *SettingsService) Save(settings models.Settings) error {
	if !settingsValid(settings) {
		return errors.New("Settings invalid")
	}
	defaults, err := s.getDefaults()
	if err != nil {
		return err
	}
	merged := s.merge(defaults, settings)
	yamlBytes, err := yaml.Marshal(merged)
	if err != nil {
		return err
	}
	err = os.WriteFile(s.filePath, yamlBytes, 0755)
	if err == nil {
		logrus.Info("Sending Settings update event")
		ev := models.NewSettingsUpdatedEvent(merged)
		s.dispatcher.Dispatch(ev)
	}
	return err
}

func (s *SettingsService) createFileIfMissing(settings models.Settings) error {
	if _, err := os.Stat(s.filePath); os.IsNotExist(err) {
		os.MkdirAll(filepath.Dir(s.filePath), 0755)
		file, err := os.Create(s.filePath)
		if err != nil {
			return fmt.Errorf("Failed to create empty config file %s", s.filePath)
		}
		defaultYaml, err := yaml.Marshal(settings)
		if err != nil {
			return fmt.Errorf(
				"Failed to write default settings into newly created settings file: %s", err.Error())
		}
		if _, err = file.Write(defaultYaml); err != nil {
			return fmt.Errorf(
				"Failed to write default settings into newly created settings file: %s", err.Error())
		}
	}
	return nil
}

func settingsValid(settings models.Settings) bool {
	return utils.IsWritableDirectory(settings.RecordingsDirectory)
}
