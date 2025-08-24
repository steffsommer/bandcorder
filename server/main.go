package main

import (
	"embed"
	"log"
	"server/internal/pkg/controllers"
	"server/internal/pkg/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"golang.org/x/net/context"
)

const API_PORT = 6000

//go:embed all:frontend/dist
var assets embed.FS

func main() {

	// Set up Web UI / wails app

	app := NewApp()
	err := wails.Run(&options.App{
		Title:  "server",
		Width:  1920,
		Height: 1080,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		Fullscreen:       true,
		OnStartup:        app.startup,
		// delay webserver startup for now to work around 
		OnDomReady:       startWebserver,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal("Failed to start wails app", err)
	}

}

func startWebserver(_ctx context.Context) {
	// Compose application, set up REST API + websockets
	go func() {
		r := gin.Default()

		recorder := services.NewRecorderService()
		recordingController := controllers.NewRecordingController(recorder)

		r.GET("/recording/start", recordingController.HandleStart)
		r.GET("/recording/stop", recordingController.HandleStop)

		log.Printf("Server starting on localhost:%d\n", API_PORT)

		if err := r.Run(":" + strconv.Itoa(API_PORT)); err != nil {
			log.Fatal("Failed to start server:", err)
		}
	}()

}
