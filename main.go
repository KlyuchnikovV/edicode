package main

import (
	"embed"
	_ "embed"

	"github.com/KlyuchnikovV/edicode/api"
	"github.com/KlyuchnikovV/edicode/core"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"golang.org/x/net/context"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	var api = api.New()

	if err := wails.Run(&options.App{
		Width:            1000,
		Height:           700,
		Title:            "edicode",
		Assets:           assets,
		BackgroundColour: options.NewRGB(255, 255, 255),
		OnStartup:        onStartup(api),
		Bind: []interface{}{
			api,
		},
	}); err != nil {
		panic(err)
	}
}

func onStartup(api *api.Api) func(ctx context.Context) {
	return func(ctx context.Context) {
		c, err := core.New(ctx, "main.go")
		if err != nil {
			panic(err)
		}

		api.Bind(c)
	}
}
