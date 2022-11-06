package main

import (
	"embed"

	"github.com/KlyuchnikovV/edicode/api"
	"github.com/KlyuchnikovV/edicode/core"
	"github.com/wailsapp/wails/v2/pkg/application"
	"github.com/wailsapp/wails/v2/pkg/options"
	"golang.org/x/net/context"
)

//go:embed frontend/dist
var assets embed.FS

func main() {
	var api = api.New()
	c, err := core.New("main.go")
	if err != nil {
		panic(err)
	}

	api.Bind(c)

	// var binds = []interface{}{
	// 	api,
	// 	&plugins.Plugin{},
	// }

	// for _, bind := range c.Manager.Plugins() {
	// 	binds = append(binds, bind)
	// }

	mainApp := application.NewWithOptions(&options.App{
		Width:            1000,
		Height:           700,
		Title:            "edicode",
		Assets:           assets,
		BackgroundColour: options.NewRGB(40, 44, 52),
		OnStartup:        onStartup(c),
		Bind: []interface{}{
			api,
			c.Manager,
		},
	})

	if err := mainApp.Run(); err != nil {
		panic(err)
	}
}

func onStartup(c *core.Core) func(ctx context.Context) {
	return func(ctx context.Context) {
		c.Init(ctx)
	}
}
