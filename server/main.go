package main

import (
	"embed"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
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
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		log.Fatal("Failed to start wails app", err)
	}

}
