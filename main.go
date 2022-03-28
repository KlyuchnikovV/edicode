package main

import (
	_ "embed"

	"github.com/KlyuchnikovV/edicode/api"
	"github.com/KlyuchnikovV/edicode/core"
	"github.com/wailsapp/wails"
	"golang.org/x/net/context"
)

//go:embed ui/public/build/bundle.js
var js string

//go:embed ui/public/index.html
var html string

//go:embed ui/public/build/bundle.css
var css string

func main() {
	app := wails.CreateApp(&wails.AppConfig{
		Width:  1000,
		Height: 700,
		Title:  "edicode",
		JS:     js,
		CSS:    css,
		HTML:   html,
		Colour: "#FFFFFF",
	})

	c, err := core.New(context.Background(), "main.go", "main_test.go")
	if err != nil {
		panic(err)
	}

	a := api.New(c)
	a.Bind(app)

	if err := c.Start(); err != nil {
		panic(err)
	}

	if err := app.Run(); err != nil {
		panic(err)
	}
}
