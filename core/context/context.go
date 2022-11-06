package context

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Context struct {
	context.Context
	cancel context.CancelFunc

	init func()
}

func New(ctx context.Context, init func()) *Context {
	return &Context{
		Context: ctx,
		init:    init,
	}
}

func (ctx *Context) Start() error {
	if ctx.cancel != nil {
		return fmt.Errorf("already started")
	}

	ctx.Context, ctx.cancel = context.WithCancel(ctx.Context)

	return nil
}

func (ctx *Context) Cancel() error {
	if ctx.cancel == nil {
		return fmt.Errorf("already cancelled")
	}

	ctx.cancel()
	ctx.cancel = nil

	return nil
}

func (ctx *Context) On(object, action string, handler func(...interface{})) {
	runtime.EventsOn(ctx.Context, object+"_"+action, handler)
}

func (ctx *Context) Emit(object, action string, data interface{}) {
	ctx.EmitTimed(object, action, time.Now().UnixNano(), data)
}

func (ctx *Context) EmitTimed(object, action string, timestamp int64, data interface{}) {
	log.Printf("HIGHLIGHT: %s changed %#v", object+"_"+action, data)

	runtime.EventsEmit(ctx.Context, object+"_"+action, DataWithTime{
		Data:      data,
		Timestamp: timestamp,
	})
}

type DataWithTime struct {
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

// func (ctx *Context) OnInit(object string, f func(...interface{})) {
// 	ctx.runtime.Events.Once(object+"_init", func(optionalData ...interface{}) {
// 		f(optionalData...)
// 	})
// }
