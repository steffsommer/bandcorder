package main

import (
	"context"
	"log"
	"server/internal/pkg/controllers"
	"server/internal/pkg/interfaces"
	"server/internal/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

const API_PORT = 6000

// App struct
type App struct {
	ctx      context.Context
	recorder interfaces.Recorder
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
	recordingController := controllers.NewRecordingController(recorder)

	//set up REST API + websockets
	go func() {
		r := gin.Default()

		r.GET("/recording/start", recordingController.HandleStart)
		r.GET("/recording/stop", recordingController.HandleStop)

		log.Printf("Server starting on localhost:%d\n", API_PORT)

		if err := r.Run(":" + strconv.Itoa(API_PORT)); err != nil {
			log.Fatal("Failed to start server:", err)
		}

	}()

	a.ctx = ctx
}

func (a *App) StartRecording() error {
	log.Println("Starting recording")
	return a.recorder.Start()
}

func (a *App) StopRecording() error {
	return a.recorder.Stop()
}

func (a *App) AbortRecording() error {
	return a.recorder.Abort()
}
