package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	goContext "context"

	"github.com/KlyuchnikovV/edicode/core/context"
	"github.com/KlyuchnikovV/edicode/core/plugins"
	"github.com/KlyuchnikovV/edicode/core/watcher"
	buffer "github.com/KlyuchnikovV/simple_buffer"
	"github.com/mitchellh/mapstructure"
)

type Core struct {
	context.Context

	buffers map[string]*buffer.Buffer

	*plugins.Manager

	watcher *watcher.Watcher

	paths []string
}

func New(paths ...string) (*Core, error) {
	var (
		core = Core{
			buffers: make(map[string]*buffer.Buffer, len(paths)),
		}
		err error
	)

	// core.Context = *context.New(ctx, core.Init)
	// core.watcher, err = watcher.New(ctx, &core, paths...)
	// if err != nil {
	// 	return nil, err
	// }
	core.paths = paths

	core.Manager, err = plugins.New(&core, "./configs/config.json")
	if err != nil {
		return nil, err
	}

	// core.On("buffer", "changed", core.on("buffer_changed", core.OnBufferChange))
	return &core, nil
}

func (core *Core) Init(ctx goContext.Context) {
	var err error

	core.Context = *context.New(ctx, nil)
	core.watcher, err = watcher.New(ctx, core, core.paths...)
	if err != nil {
		panic(err)
	}

	if err := core.Manager.LoadPlugins(); err != nil {
		panic(err)
	}
}

func (core *Core) on(eventName string, handler plugins.On) func(...interface{}) {
	return func(args ...interface{}) {
		log.Printf("PLUGIN: event %s, args %#v", eventName, args)
		var event plugins.Event
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

		event = plugins.BufferChangeEvent{
			Buffer:    data.Data.(string),
			Timestamp: data.Timestamp,
		}
		if err := handler(event); err != nil {
			panic(err) // TODO: errors channel in 'p'
		}
	}
}

func (core *Core) OpenFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	core.buffers[path] = buffer.NewFromBytes(core.Context, path, bytes...)
	core.Emit("buffer", "opened", path)

	return nil
}

func (core *Core) CloseFile(path string) error {
	if _, ok := core.buffers[path]; !ok {
		return fmt.Errorf("buffer not found")
	}

	delete(core.buffers, path)
	core.Emit("buffer", "closed", path)

	return nil
}

func (core *Core) ReloadFile(path string) error {
	core.Emit("buffer", "changed", path)
	return nil
}

func (core *Core) GetFileNames(path string) ([]string, error) {
	panic("implement me")
}

func (core *Core) GetPluginsPath() string {
	return "./plugins"
}
