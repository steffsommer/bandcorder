package main

import (
	"context"
	"embed"
	"log"
	"os"
	"path/filepath"
	"server/internal/pkg/controllers"
	"server/internal/pkg/facades"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/models"
	"server/internal/pkg/services"
	"server/internal/pkg/services/cyclic_sender"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

const API_PORT = 6000
const AUDIO_CHANNEL_COUNT = 1 // Mono
const SAMPLE_RATE_HZ = 44100
const SETTINGS_FOLDER_NAME = "bandcorder"
const SETTINGS_FILE_NAME = "config.yaml"

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Failed to determine user config directory: " + err.Error())
	}
	settingsFilePath := filepath.Join(configDir, SETTINGS_FOLDER_NAME, SETTINGS_FILE_NAME)

	// NewApp creates a new App application struct
	settingsService := services.NewSettingsService(settingsFilePath)
	settings, err := settingsService.Load()
	if err != nil {
		logrus.Fatalf("Failed to load settings: %s", err.Error())
	}
	websocketController := controllers.NewWebsocketController()
	uiSenderService := services.NewUiSenderService()
	broadcastSender := services.NewBroadcastSender(
		[]interfaces.EventDispatcher{
			websocketController,
			uiSenderService,
		},
	)

	eventbus := cyclic_sender.NewCyclicSender(broadcastSender)

	storageService := services.NewFileSystemStorageService(
		settings.RecordingsDirectory,
		AUDIO_CHANNEL_COUNT,
		SAMPLE_RATE_HZ,
	)
	processor := services.NewAudioProcessorService(broadcastSender)
	recorder := services.NewRecorderService(storageService, processor)
	recordingFacade := facades.NewRecordingFacade(eventbus, recorder)
	recordingController := controllers.NewRecordingController(recordingFacade)

	r := gin.Default()

	r.POST("/recording/start", recordingController.HandleStart)
	r.POST("/recording/stop", recordingController.HandleStop)
	r.POST("/recording/abort", recordingController.HandleAbort)
	r.GET("/ws", websocketController.HandleWebsocketUpgrade)

	modelExporter := models.ModelExporter{}
	err = wails.Run(&options.App{
		Title:  "server",
		Width:  1920,
		Height: 1080,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		WindowStartState: options.Maximised,
		OnStartup: func(ctx context.Context) {
			if err := recorder.Init(); err != nil {
				panic("Failed to init recorder service: " + err.Error())
			}
			uiSenderService.Init(ctx)
			eventbus.StartSendingPeriodicUpdates()
			go func() {
				log.Printf("Server starting on localhost:%d\n", API_PORT)

				if err := r.Run(":" + strconv.Itoa(API_PORT)); err != nil {
					log.Fatal("Failed to start server:", err)
				}

			}()
		},
		Bind: []interface{}{
			&modelExporter,
			recordingFacade,
			storageService,
			settingsService,
		},
	})

	if err != nil {
		log.Fatal("Failed to start wails app", err)
	}

}
