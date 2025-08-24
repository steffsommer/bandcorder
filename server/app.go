package main

import (
	"context"
	"fmt"
	"log"
	"server/internal/pkg/controllers"
	"server/internal/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

const API_PORT = 6000

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	fmt.Println("Startup handler running")
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
