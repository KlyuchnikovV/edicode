package plugin

import (
	"fmt"
	"log"
	"plugin"

	"github.com/KlyuchnikovV/edicode/core/context"
	"github.com/KlyuchnikovV/edicode/types"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
)

type (
	Init func(types.Core) error
	On   func(Event) error
	Bind func() map[string]On

	Event interface {
		// Source() string
		// Destination() string
		// String() string
	}

	PluginInterface interface {
		Init(types.Core) error
		Bind() map[string]On
	}

	Plugin struct {
		name   string
		path   string
		plugin plugin.Plugin

		pluginInterface PluginInterface
	}
)

func New(c types.Core, name, path string) (*Plugin, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, err
	}

	pl, err := p.Lookup("Plugin")
	if err != nil {
		return nil, err
	}

	typedPlugin, ok := pl.(PluginInterface)
	if !ok {
		return nil, fmt.Errorf("plugin has wrong type") //TODO: errors package
	}

	if err := typedPlugin.Init(c); err != nil {
		return nil, err
	}

	return &Plugin{
		name:            name,
		path:            path,
		plugin:          *p,
		pluginInterface: typedPlugin,
	}, nil
}

func (p *Plugin) WailsInit(runtime *wails.Runtime) error {
	for event, handler := range p.pluginInterface.Bind() {
		runtime.Events.On(
			event, p.on(event, handler),
		)
	}

	return nil
}

func (p *Plugin) on(event string, handler On) func(...interface{}) {
	return func(args ...interface{}) {
		log.Printf("PLUGIN: event %s, args %#v", event, args)
		var event Event
		if len(args) == 0 {
			panic("args is zero") // TODO: errors channel in 'p'
		}

		var (
			data context.DataWithTime
			ok   bool
		)
		switch typed := args[0].(type) {
		case map[string]interface{}:
			if err := mapstructure.Decode(typed, &data); err != nil {
				panic(err)
			}
		default:
			data, ok = typed.(context.DataWithTime)
			if !ok {
				panic("!ok")
			}
		}

		event = BufferChangeEvent{
			Buffer:    data.Data.(string),
			Timestamp: data.Timestamp,
		}
		if err := handler(event); err != nil {
			panic(err) // TODO: errors channel in 'p'
		}
	}
}

type BufferChangeEvent struct {
	Buffer    string
	Timestamp int64
}
