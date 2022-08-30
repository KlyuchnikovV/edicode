package api

import (
	"log"

	"github.com/KlyuchnikovV/edicode/core"
	"github.com/KlyuchnikovV/edicode/types"
	"github.com/KlyuchnikovV/simple_buffer"
	"github.com/mitchellh/mapstructure"
)

type Api struct {
	core *core.Core
}

func New() *Api {
	return &Api{}
}

func (api *Api) Bind(core *core.Core) {
	api.core = core
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
	var event simple_buffer.KeyboardEvent
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

func (api *Api) LengthOfLine(request types.GetLineLengthRequest) (int, error) {
	return api.core.LengthOfLine(request)
}

func (api *Api) LengthOfBuffer(buffer string) (int, error) {
	return api.core.LengthOfBuffer(buffer)
}

func (api *Api) GetCursor(buffer string) (*types.GetCaretResponse, error) {
	return api.core.GetCursor(buffer)
}

func (api *Api) SetCursor(data map[string]interface{}) (*types.GetCaretResponse, error) {
	var event types.CaretMovedEvent
	if err := mapstructure.Decode(data, &event); err != nil {
		return nil, err
	}
	return api.core.SetCursor(event)
}
