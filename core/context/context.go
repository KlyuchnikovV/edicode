package context

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wailsapp/wails"
)

type Context struct {
	ctx    context.Context
	cancel context.CancelFunc

	runtime *wails.Runtime

	init func()
}

func New(ctx context.Context, init func()) *Context {
	return &Context{
		ctx:  ctx,
		init: init,
	}
}

func (ctx *Context) WailsInit(runtime *wails.Runtime) error {
	ctx.runtime = runtime

	if ctx.init != nil {
		ctx.init()
	}

	return nil
}

func (ctx *Context) Start() error {
	if ctx.cancel != nil {
		return fmt.Errorf("already started")
	}

	ctx.ctx, ctx.cancel = context.WithCancel(ctx.ctx)

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
	ctx.runtime.Events.On(object+"_"+action, handler)
}

func (ctx *Context) Emit(object, action string, data interface{}) {
	ctx.EmitTimed(object, action, time.Now().UnixNano(), data)
}

func (ctx *Context) EmitTimed(object, action string, timestamp int64, data interface{}) {
	log.Printf("HIGHLIGHT: %s changed %#v", object+"_"+action, data)
	ctx.runtime.Events.Emit(object+"_"+action, DataWithTime{
		Data:      data,
		Timestamp: timestamp,
	})
}

func (ctx *Context) Done() <-chan struct{} {
	return ctx.ctx.Done()
}

type DataWithTime struct {
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}

func (ctx *Context) OnInit(object string, f func(...interface{})) {
	ctx.runtime.Events.Once(object+"_init", func(optionalData ...interface{}) {
		f(optionalData...)
	})
}
