package plugins

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/KlyuchnikovV/edicode/types"
	buffer "github.com/KlyuchnikovV/simple_buffer"
)

type Context interface {
	OpenFile(path string) error
	CloseFile(path string) error
	ReloadFile(path string) error
	Buffers() map[string]*buffer.Buffer
	GetBuffer(name string) (*types.BufferData, error)
	GetBufferNames() []string
	LengthOfLine(request types.GetLineLengthRequest) (int, error)
	LengthOfBuffer(buffer string) (int, error)
	GetCursor(buffer string) (*types.GetCaretResponse, error)
	SaveBuffer(name string) error
	GetFileNames(string) ([]string, error)
	GetPluginsPath() string
	// Events
	On(object, action string, handler func(...interface{}))
	Emit(object, action string, data interface{})
	EmitTimed(object, action string, timestamp int64, data interface{})
	// TODO:
	GetActionsList(filterBy string) []Action
}

type (
	BackendPlugin interface {
		Init(Context) error
		Functions() map[string]func(interface{}) (interface{}, error)
	}
	ActionsPlugin interface {
		Actions() map[string]func() error
	}
)
type (
	Init func(Context) error
	On   func(Event) error
	Bind func() map[string]On

	Event interface {
		// Source() string
		// Destination() string
		// String() string
	}
)

type BufferChangeEvent struct {
	Buffer    string
	Timestamp int64
}

type Config struct {
	BuildPath   string      `json:"buildPath"`
	PluginsPath string      `json:"pluginsPath"`
	Plugins     PluginNames `json:"plugins"`
}

type PluginNames []string

func (config *Config) Load(path string) (err error) {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer func() {
		e := file.Close()

		if err == nil {
			err = e
		} else {
			err = fmt.Errorf("%s: %w", err, e)
		}
	}()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, config)
}

func (names PluginNames) Has(name string) bool {
	for _, plugin := range names {
		if plugin == name {
			return true
		}
	}

	return false
}

type PanelEntry struct {
	Name string `json:"name"`
	File string `json:"file"`
	Icon string `json:"icon"`
}

func (entry *PanelEntry) Cleanup() {
	entry.Name = strings.ToLower(entry.Name)
	entry.File = strings.TrimSuffix(strings.Trim(entry.File, "./"), ".svelte")
	entry.Icon = strings.ToLower(entry.Icon)
}

type Action struct {
	Name   string `json:"name"`
	Action string `json:"action"`
}

type PluginConfig struct {
	Name   string `json:"name"`
	Panels Panels `json:"panels"`
}

type Panels struct {
	RightPanel *PanelEntry `json:"right"`
}

func (config *PluginConfig) Load(directory string) error {
	file, err := os.Open(
		path.Join(directory, "plugin.json"),
	)
	if err != nil {
		return err
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, config); err != nil {
		return err
	}

	defer config.Cleanup()

	return config.Validate()
}

func (config *PluginConfig) Cleanup() {
	config.Panels.RightPanel.Cleanup()
}

func (config *PluginConfig) Validate() error {
	return nil
}
