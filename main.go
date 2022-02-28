package main

import (
	_ "embed"

	"github.com/KlyuchnikovV/edigode/api"
	"github.com/KlyuchnikovV/edigode/core"
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
		Title:  "edigode",
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
