package core

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/KlyuchnikovV/buffer"
	"github.com/KlyuchnikovV/edigode/core/plugin"
	"github.com/KlyuchnikovV/edigode/core/syntax"
	"github.com/KlyuchnikovV/edigode/types"
)

func (core *Core) GetBuffer(name string) (*types.BufferData, error) {
	buf, ok := core.buffers[name]
	if !ok {
		return nil, fmt.Errorf("buffer '%s' not found", name)
	}

	lines, err := buf.GetLines()
	if err != nil {
		return nil, err
	}

	result, err := syntax.New().Tokenize(strings.Join(lines, "\n"))
	if err != nil {
		return nil, err
	}

	return &types.BufferData{
		Buffer:    name,
		Lines:     result,
		Timestamp: time.Now().UnixNano(),
	}, nil
}

func (core *Core) HandleKeyboardEvent(event types.KeyboardEvent) error {
	buf, ok := core.buffers[event.Buffer]
	if !ok {
		return fmt.Errorf("buf '%s' not found", event.Buffer)
	}

	switch {
	case event.Key == "Backspace":
		log.Printf("HANDLE: backspace event is %#v", event)
		buf.KeyboardEvents() <- event
	case event.Key == "ArrowLeft":
		buf.KeyboardEvents() <- event
	case event.Key == "ArrowRight":
		buf.KeyboardEvents() <- event
	case event.Key == "ArrowUp":
		buf.KeyboardEvents() <- event
	case event.Key == "ArrowDown":
		buf.KeyboardEvents() <- event
	case event.Key == "Enter":
		buf.KeyboardEvents() <- event
	case event.Key == "Escape":
		log.Printf("EXITING")
		os.Exit(0)
	case len(event.Key) == 1:
		log.Printf("HANDLE: event is %#v", event)
		buf.KeyboardEvents() <- event
	case event.Key == "Tab":
		buf.KeyboardEvents() <- event
	}

	return nil
}

func (core *Core) Buffers() map[string]*buffer.Buffer {
	return core.buffers
}

func (core *Core) GetBufferNames() []string {
	var result = make([]string, 0, len(core.buffers))
	for name := range core.buffers {
		result = append(result, name)
	}

	sort.Slice(result, func(i, j int) bool {
		for k := 0; k < len(result[i]); k++ {
			if k >= len(result[j]) {
				return false
			}
			if result[i][k] != result[j][k] {
				return result[i][k] < result[j][k]
			}
		}
		return true
	})

	log.Printf("LOG: buffers are %#v", result)
	return result
}

func (core *Core) OnBufferChange(event plugin.Event) error {
	log.Printf("HIGHLIGHT: got event %#v", event)
	buffer, err := core.GetBuffer(event.(plugin.BufferChangeEvent).Buffer)
	if err != nil {
		return err
	}

	for i := range buffer.Lines {
		for j := range buffer.Lines[i] {
			buffer.Lines[i][j].Classes = nil
		}
	}

	// TODO: highlight

	// send data to backdrop
	core.EmitTimed("highlight", "changed", buffer.Timestamp, buffer)

	return nil
}
