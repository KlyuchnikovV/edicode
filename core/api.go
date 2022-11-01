package core

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/KlyuchnikovV/edicode/core/plugins"
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

	// sort.Slice(result, func(i, j int) bool {
	// 	for k := 0; k < len(result[i]); k++ {
	// 		if k >= len(result[j]) {
	// 			return false
	// 		}
	// 		if result[i][k] != result[j][k] {
	// 			return result[i][k] < result[j][k]
	// 		}
	// 	}
	// 	return true
	// })

	log.Printf("LOG: buffers are %#v", result)
	return result
}

func (core *Core) OnBufferChange(event plugins.Event) error {
	log.Printf("HIGHLIGHT: got event %#v", event)
	buffer, err := core.GetBuffer(event.(plugins.BufferChangeEvent).Buffer)
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

	buf.SetStartAndEnd(event.Start, event.End)

	start, end := buf.GetSelection()
	core.Emit("cursor_moved", event.Buffer, types.GetCaretResponse{
		Buffer: event.Buffer,
		Start:  start,
		End:    end,
	})

	return core.GetCursor(event.Buffer)
}

func (core *Core) MouseUp(event types.MouseEvent) error {
	buf, ok := core.buffers[event.Buffer]
	if !ok {
		return fmt.Errorf("buf '%s' not found", event.Buffer)
	}

	buf.SetEnd(event.Line, event.Offset)

	return nil
}

func (core *Core) MouseDown(event types.MouseEvent) error {
	buf, ok := core.buffers[event.Buffer]
	if !ok {
		return fmt.Errorf("buf '%s' not found", event.Buffer)
	}

	buf.SetStart(event.Line, event.Offset)

	fmt.Println(event)
	return nil
}

func (core *Core) SaveBuffer(name string) error {
	buf, ok := core.buffers[name]
	if !ok {
		return fmt.Errorf("buf '%s' not found", name)
	}

	file, err := os.OpenFile(buf.Name(), os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return err
	}

	_, err = file.WriteAt([]byte(buf.String()), 0)
	return err
}

func (core *Core) GetActionsList(filterBy string) []plugins.Action {
	var result = make([]plugins.Action, 0)

	for _, plug := range core.Manager.Plugins() {
		if plug.Actions == nil {
			continue
		}

		for name := range plug.Actions() {
			result = append(result, plugins.Action{
				Name:   fmt.Sprintf("%s: %s", plug.Config.Name, name),
				Action: fmt.Sprintf("%s.%s", strings.ToLower(plug.Config.Name), name),
			})
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result
}
