package main

import (
	recorder "bandcorder/internal/pkg"
	"embed"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	logrus.Info("Starting bandcorder")
	recorder, err := recorder.NewRecorder()
	if err != nil {
		logrus.Errorf("Failed to create recorder instance: %s", err.Error())
	}
	if err := recorder.Start(); err != nil {
		logrus.Errorf("Failed to start recording: %s", err.Error())
	}
	time.Sleep(5 * time.Second)
	if err := recorder.Stop(); err != nil {
		logrus.Errorf("Failed to stop recording: %s", err.Error())
	}

	// Create application with options
	err = wails.Run(&options.App{
		Title:      "bandcorder",
		Fullscreen: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
