package core

import (
	"io/ioutil"
	"log"
	"os"

	goContext "context"

	"github.com/KlyuchnikovV/edicode/core/actions"
	"github.com/KlyuchnikovV/edicode/core/context"
	"github.com/KlyuchnikovV/edicode/core/plugin"
	buffer "github.com/KlyuchnikovV/simple_buffer"
	"github.com/mitchellh/mapstructure"
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
	actions []actions.Action
}

func New(ctx goContext.Context, paths ...string) (*Core, error) {
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

		core.buffers[path] = buffer.NewFromBytes(core.Context, path, bytes...)
	}

	core.actions = []actions.Action{
		actions.NewSaveFile(core.SaveBuffer),
		&actions.OpenFile{},
	}

	// for _, pl := range core.plugins {
	// 	app.Bind(&pl)
	// }

	// p, err := plugin.New(
	// 	&core,
	// 	fmt.Sprintf("highlighter_%s", "go"),
	// 	"../highlight/build/highlight.so",
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// core.plugins = []plugin.Plugin{}

	core.On("buffer", "changed", core.on("buffer_changed", core.OnBufferChange))
	return &core, nil
}

func (core *Core) Init() {
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
