package main

import (
	"context"
	"embed"
	"log"
	"server/internal/pkg/controllers"
	"server/internal/pkg/facades"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/services"
	"server/internal/pkg/services/eventbus"
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

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// NewApp creates a new App application struct
	settings, err := services.NewSettingsService().Load()
	if err != nil {
		logrus.Fatalf("Failed to load settings: %w", err)
	}
	websocketController := controllers.NewWebsocketController()
	uiSenderService := services.NewUiSenderService()
	broadcastSender := services.NewBroadcastSender(
		[]interfaces.Sender{
			websocketController,
			uiSenderService,
		},
	)

	eventbus := eventbus.NewEventBus(broadcastSender)
	storageService := services.NewFileSystemStorageService(
		settings.RecordingsDirectory,
		AUDIO_CHANNEL_COUNT,
		SAMPLE_RATE_HZ,
	)
	recorder := services.NewRecorderService(storageService)
	recordingFacade := facades.NewRecordingFacade(eventbus, recorder)
	recordingController := controllers.NewRecordingController(recordingFacade)

	r := gin.Default()

	r.POST("/recording/start", recordingController.HandleStart)
	r.POST("/recording/stop", recordingController.HandleStop)
	r.POST("/recording/abort", recordingController.HandleAbort)
	r.POST("/ws", websocketController.HandleWebsocketUpgrade)

	err = wails.Run(&options.App{
		Title:  "server",
		Width:  1920,
		Height: 1080,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 0},
		Fullscreen:       true,
		OnStartup: func(ctx context.Context) {
			if err := recorder.Init(); err != nil {
				panic("Failed to init recorder service: " + err.Error())
			}
			uiSenderService.Init(ctx)
			go func() {
				log.Printf("Server starting on localhost:%d\n", API_PORT)

				if err := r.Run(":" + strconv.Itoa(API_PORT)); err != nil {
					log.Fatal("Failed to start server:", err)
				}

			}()
		},
		OnDomReady: func(_ context.Context) {
			eventbus.StartSendingPeriodicUpdates()
		},
		Bind: []interface{}{
			recordingFacade,
		},
	})

	if err != nil {
		log.Fatal("Failed to start wails app", err)
	}

}
