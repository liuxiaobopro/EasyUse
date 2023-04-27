package main

import (
	"embed"
	"flag"

	"EasyUse/server"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

type conf struct {
	Runmode string
}

func main() {
	var c = new(conf)
	flag.StringVar(&c.Runmode, "runmode", "", "runmode")
	flag.Parse()

	if c.Runmode == "server" {
		server.Start()
	} else {
		// server start
		go server.Start()

		// Create an instance of the app structure
		app := NewApp()

		// Create application with options
		err := wails.Run(&options.App{
			Title:  "易用 EasyUse",
			Width:  1024,
			Height: 768,
			AssetServer: &assetserver.Options{
				Assets: assets,
			},
			// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
			OnStartup: app.startup,
			Bind: []interface{}{
				app,
			},
		})

		if err != nil {
			println("Error:", err.Error())
		}
	}
}
