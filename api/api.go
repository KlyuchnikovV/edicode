package api

import (
	"log"

	"github.com/KlyuchnikovV/edigode/core"
	"github.com/KlyuchnikovV/edigode/types"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
)

type Api struct {
	core *core.Core
}

func New(c *core.Core) *Api {
	return &Api{
		core: c,
	}
}

func (api *Api) Bind(app *wails.App) {
	app.Bind(api)
	api.core.Bind(app)
}

func (api *Api) GetBuffer(name string) (types.BufferData, error) {
	buffer, err := api.core.GetBuffer(name)
	if err != nil || buffer == nil {
		return types.BufferData{}, err
	}
	return *buffer, nil
}

func (api *Api) HandleKeyboardEvent(data map[string]interface{}) error {
	log.Printf("API: %#v", data)
	var event types.KeyboardEvent
	if err := mapstructure.Decode(data, &event); err != nil {
		return err
	}
	err := api.core.HandleKeyboardEvent(event)
	if err != nil {
		log.Printf("ERROR: %#v", err)
	}
	return err
}

func (api *Api) GetBufferNames() []string {
	return api.core.GetBufferNames()
}
