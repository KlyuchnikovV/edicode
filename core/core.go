package core

import (
	ctx "context"
	"io/ioutil"
	"log"
	"os"

	"github.com/KlyuchnikovV/edicode/core/context"
	"github.com/KlyuchnikovV/edicode/core/plugin"
	buffer "github.com/KlyuchnikovV/simple_buffer"
	"github.com/mitchellh/mapstructure"
	"github.com/wailsapp/wails"
)

/* TODO:
- delete, copy, paste selection
- undo, redo
-

*/

type Core struct {
	context.Context

	buffers map[string]*buffer.Buffer

	plugins []plugin.Plugin
}

func New(ctx ctx.Context, paths ...string) (*Core, error) {
	var core = Core{
		buffers: make(map[string]*buffer.Buffer, len(paths)),
	}
	core.Context = *context.New(ctx, core.Init)

	for _, path := range paths {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}

		core.buffers[path] = buffer.NewFromBytes(bytes...)
	}

	// p, err := plugin.New(
	// 	&core,
	// 	fmt.Sprintf("highlighter_%s", "go"),
	// 	"../highlight/build/highlight.so",
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// core.plugins = []plugin.Plugin{}

	return &core, nil
}

func (core *Core) Init() {
	core.On("buffer", "changed", core.on("buffer_changed", core.OnBufferChange))
}

func (core *Core) Bind(app *wails.App) {
	app.Bind(&core.Context)

	// for _, buf := range core.Buffers() {
	// 	app.Bind(buf)
	// }

	for _, pl := range core.plugins {
		app.Bind(&pl)
	}
}

func (core *Core) Start() error {
	// for _, buf := range core.buffers {
	// 	if err := buf.Start(); err != nil {
	// 		return nil
	// 	}
	// }

	return nil
}

func (core *Core) Cancel() error {
	// for _, buf := range core.buffers {
	// 	if err := buf.Cancel(); err != nil {
	// 		return nil
	// 	}
	// }

	return nil
}

func (core *Core) on(eventName string, handler plugin.On) func(...interface{}) {
	return func(args ...interface{}) {
		log.Printf("PLUGIN: event %s, args %#v", eventName, args)
		var event plugin.Event
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

		event = plugin.BufferChangeEvent{
			Buffer:    data.Data.(string),
			Timestamp: data.Timestamp,
		}
		if err := handler(event); err != nil {
			panic(err) // TODO: errors channel in 'p'
		}
	}
}
