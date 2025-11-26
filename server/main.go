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

	var storageService *services.FileSystemStorageService

	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Failed to determine user config directory: " + err.Error())
	}
	settingsFilePath := filepath.Join(configDir, SETTINGS_FOLDER_NAME, SETTINGS_FILE_NAME)

	websocketController := controllers.NewWebsocketController()
	uiSenderService := services.NewUiEventBus()
	broadcastSender := services.NewBroadcastSender(
		[]interfaces.EventDispatcher{
			websocketController,
			uiSenderService,
		},
	)

	// NewApp creates a new App application struct
	settingsService := services.NewSettingsService(settingsFilePath, uiSenderService)
	settings, err := settingsService.Load()
	if err != nil {
		logrus.Fatalf("Failed to load settings: %s", err.Error())
	}

	cyclicSender := services.NewCyclicRecordingEventSender(broadcastSender)
	timeProvider := services.NewRealTimeProvider()

	storageService = services.NewFileSystemStorageService(
		settings.RecordingsDirectory,
		AUDIO_CHANNEL_COUNT,
		SAMPLE_RATE_HZ,
		timeProvider,
		broadcastSender,
		uiSenderService,
	)
	playbackService := services.NewAudioPlaybackService()
	storageFacade := facades.NewFileSystemStorageFacade(playbackService, storageService)
	processor := services.NewAudioProcessorService(broadcastSender)
	recorder := services.NewRecorderService(storageFacade, processor)
	recordingFacade := facades.NewRecordingFacade(cyclicSender, recorder, playbackService)
	recordingController := controllers.NewRecordingController(recordingFacade)

	fileController := controllers.NewFileController(storageFacade)
	metronomeService := services.NewMetronomeService(86, broadcastSender, playbackService)

	metronomeController := controllers.NewMetronomeController(metronomeService)

	r := gin.Default()

	r.POST("/recording/start", recordingController.HandleStart)
	r.POST("/recording/stop", recordingController.HandleStop)
	r.POST("/recording/abort", recordingController.HandleAbort)

	r.POST("/metronome/start", metronomeController.HandleStart)
	r.POST("/metronome/stop", metronomeController.HandleStop)
	r.POST("/metronome/updateBpm", metronomeController.HandleUpdateBpm)
	r.GET("/metronome/state", metronomeController.HandleGetState)

	r.POST("/files/renameLast", fileController.HandleRenameLast)
	r.GET("/ws", websocketController.HandleWebsocketUpgrade)

	ipService := &services.IPService{}
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
			if err := playbackService.Init(); err != nil {
				panic("Failed to init playback service: " + err.Error())
			}
			if err := recorder.Init(); err != nil {
				panic("Failed to init recorder service: " + err.Error())
			}
			uiSenderService.Init(ctx)
			storageService.InitSubscriptions()
			cyclicSender.StartSendingPeriodicUpdates()
			go func() {
				log.Printf("Server starting on localhost:%d\n", API_PORT)

				if err := r.Run("0.0.0.0:" + strconv.Itoa(API_PORT)); err != nil {
					log.Fatal("Failed to start server:", err)
				}

			}()
		},
		Bind: []interface{}{
			&modelExporter,
			recordingFacade,
			storageFacade,
			settingsService,
			ipService,
			metronomeService,
		},
	})

	if err != nil {
		log.Fatal("Failed to start wails app", err)
	}

}
