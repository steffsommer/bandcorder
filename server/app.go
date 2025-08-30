package main

import (
	"context"
	"log"
	"server/internal/pkg/controllers"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/services"
	"server/internal/pkg/services/notifier"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const API_PORT = 6000

// App struct
type App struct {
	ctx      context.Context
	recorder interfaces.Recorder
	notifier *notifier.Notifier
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		recorder: services.NewRecorderService(),
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	recorder := services.NewRecorderService()
	if err := recorder.Init(); err != nil {
		panic("Failed to init recorder service: " + err.Error())
	}

	websocketController := controllers.NewWebsocketController()
	uiSender := services.NewUiSenderService(ctx)
	broadcastSender := services.NewBroadcastSender(
		[]interfaces.Sender{
			websocketController,
			uiSender,
		},
	)
	a.notifier = notifier.NewNotifier(broadcastSender)

	recordingController := controllers.NewRecordingController(recorder, a.notifier)

	//set up REST API + websockets
	go func() {
		r := gin.Default()

		r.POST("/recording/start", recordingController.HandleStart)
		r.POST("/recording/stop", recordingController.HandleStop)
		r.POST("/recording/abort", recordingController.HandleAbort)
		r.POST("/ws", websocketController.HandleWebsocketUpgrade)

		log.Printf("Server starting on localhost:%d\n", API_PORT)

		if err := r.Run(":" + strconv.Itoa(API_PORT)); err != nil {
			log.Fatal("Failed to start server:", err)
		}

	}()

	a.ctx = ctx
}

func (a *App) domReady(_ context.Context) {
	time.Sleep(500 * time.Millisecond)
	a.notifier.StartSendingPeriodicUpdates()
}

func (a *App) StartRecording() error {
	log.Println("Starting recording")
	err := a.recorder.Start()
	return err
}

func (a *App) StopRecording() error {
	return a.recorder.Stop()
}

func (a *App) AbortRecording() error {
	return a.recorder.Abort()
}
