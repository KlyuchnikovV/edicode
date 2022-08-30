package core

import (
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/KlyuchnikovV/edicode/core/plugin"
	"github.com/KlyuchnikovV/edicode/core/syntax"
	"github.com/KlyuchnikovV/edicode/types"
	buffer "github.com/KlyuchnikovV/simple_buffer"
)

func (core *Core) GetBuffer(name string) (*types.BufferData, error) {
	buf, ok := core.buffers[name]
	if !ok {
		return nil, fmt.Errorf("buffer '%s' not found", name)
	}

	result, err := syntax.New().TokenizeLines(buf.String())
	if err != nil {
		return nil, err
	}

	return &types.BufferData{
		Buffer:    name,
		Lines:     result,
		Timestamp: time.Now().UnixNano(),
	}, nil
}

func (core *Core) HandleKeyboardEvent(event buffer.KeyboardEvent) error {
	buffer, ok := core.buffers[event.Buffer]
	if !ok {
		return fmt.Errorf("buf '%s' not found", event.Buffer)
	}

	buffer.KeyEvents <- event

	// switch {
	// case event.Key == "Backspace":
	// 	log.Printf("HANDLE: backspace event is %#v", event)
	// 	if err := buf.Delete(1); err != nil {
	// 		panic(err)
	// 	}
	// 	// buf.KeyboardEvents() <- event
	// case event.Key == "ArrowLeft":
	// 	buf.CursorLeft(event.Shift)
	// 	// buf.KeyboardEvents() <- event
	// case event.Key == "ArrowRight":
	// 	buf.CursorRight(event.Shift)
	// 	// buf.KeyboardEvents() <- event
	// case event.Key == "ArrowUp":
	// 	buf.CursorUp(event.Shift)
	// 	// buf.KeyboardEvents() <- event
	// case event.Key == "ArrowDown":
	// 	buf.CursorDown(event.Shift)
	// 	// buf.KeyboardEvents() <- event
	// case event.Key == "Enter":
	// 	// buf.KeyboardEvents() <- event
	// 	if err := buf.Append('\n'); err != nil {
	// 		panic(err)
	// 	}
	// case event.Key == "Escape":
	// 	log.Printf("EXITING")
	// 	os.Exit(0)
	// case len(event.Key) == 1:
	// 	log.Printf("HANDLE: event is %#v", event)
	// 	if err := buf.Append(rune(event.Key[0])); err != nil {
	// 		panic(err)
	// 	}
	// 	// buf.KeyboardEvents() <- event
	// case event.Key == "Tab":
	// 	if err := buf.Append('\t'); err != nil {
	// 		panic(err)
	// 	}
	// 	// buf.KeyboardEvents() <- event
	// }

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
	core.EmitTimed("highlight_changed", buffer.Buffer, buffer.Timestamp, buffer)

	return nil
}

func (core *Core) LengthOfLine(request types.GetLineLengthRequest) (int, error) {
	buf, ok := core.buffers[request.Buffer]
	if !ok {
		return -1, fmt.Errorf("buf '%s' not found", request.Buffer)
	}

	return buf.LengthOfLine(request.Line)
}

func (core *Core) LengthOfBuffer(buffer string) (int, error) {
	buf, ok := core.buffers[buffer]
	if !ok {
		return -1, fmt.Errorf("buf '%s' not found", buffer)
	}

	return len(buf.String()), nil
}

func (core *Core) GetCursor(buffer string) (*types.GetCaretResponse, error) {
	buf, ok := core.buffers[buffer]
	if !ok {
		return nil, fmt.Errorf("buf '%s' not found", buffer)
	}

	start, end := buf.GetSelection()
	return &types.GetCaretResponse{
		Start: start,
		End:   end,
	}, nil
}

func (core *Core) SetCursor(event types.CaretMovedEvent) (*types.GetCaretResponse, error) {
	buf, ok := core.buffers[event.Buffer]
	if !ok {
		return nil, fmt.Errorf("buf '%s' not found", event.Buffer)
	}

	buf.SetSelection(event.Start, event.End)

	start, end := buf.GetSelection()
	core.Emit("cursor_moved", event.Buffer, types.GetCaretResponse{
		Buffer: event.Buffer,
		Start:  start,
		End:    end,
	})

	return core.GetCursor(event.Buffer)
}
