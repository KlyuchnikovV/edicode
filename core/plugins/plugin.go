package plugins

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"plugin"
)

type Plugin struct {
	pluginsDir string
	// buildDir   string
	name string

	// path   string
	plugin plugin.Symbol
	Config PluginConfig

	Actions   func() map[string]func() error
	Functions func() map[string]func(interface{}) (interface{}, error)
}

func NewPlugin(core Context, pluginsDir, buildDir, name string) (*Plugin, error) {
	var (
		plugin = Plugin{
			pluginsDir: pluginsDir,
			name:       name,
		}
	)
	// TODO: check plugin's hash

	// if hash is different -> build new version
	if err := plugin.BuildBackend(); err != nil {
		return nil, err
	}

	// if err := plugin.BuildFrontend(); err != nil {
	// 	return nil, err
	// }

	if err := plugin.LoadLib(core); err != nil {
		return nil, err
	}

	if err := plugin.Config.Load(plugin.Directory()); err != nil {
		return nil, err
	}

	if casted, ok := plugin.plugin.(ActionsPlugin); ok {
		plugin.Actions = casted.Actions
	}

	core.Emit("plugin_loaded", name, nil)
	return &plugin, nil
}

func (plug *Plugin) Directory() string {
	return path.Join(plug.pluginsDir, plug.name)
}

func (plug *Plugin) Executable() string {
	return fmt.Sprintf("%s.so", plug.name)
}

func (plug *Plugin) BuildBackend() error {
	var (
		cmd = exec.Command(
			"go", "build",
			`-tags=dev`,
			"-gcflags", "all=-N -l",
			`-buildmode=plugin`,
			fmt.Sprintf("-o=%s", plug.Executable()),
		)
		outBuffer, errBuffer bytes.Buffer
	)

	// Manually set the command line arguments so they are not escaped
	cmd.Dir = plug.Directory()
	cmd.Env = []string{
		`CGO_CFLAGS=-mmacosx-version-min=10.13`,
		`CGO_CXXFLAGS=-I/Users/klyuchnikovv/go/src/github.com/KlyuchnikovV/archive/edicode/build`,
		`CGO_ENABLED=1`,
		`CGO_LDFLAGS=-framework UniformTypeIdentifiers -mmacosx-version-min=10.13`,
		`GOOS=darwin`,
		`GOARCH=arm64`,
	}
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Stdout = &outBuffer
	cmd.Stderr = &errBuffer

	err := cmd.Run()
	if err == nil {
		return nil
	}

	if outBuffer.Len() > 0 {
		fmt.Printf("build plugin: %s", outBuffer.String())
	}

	return fmt.Errorf(
		"can't build plugin: %s: %s", errBuffer.String(), err,
	)
}

// func (plug *Plugin) BuildFrontend() error {
// 	var (
// 		cmd = exec.Command(
// 			"vite", "build",
// 		)
// 		outBuffer, errBuffer bytes.Buffer
// 	)

// 	// Manually set the command line arguments so they are not escaped
// 	cmd.Dir = plug.Directory()
// 	// cmd.Env = []string{
// 	// 	`CGO_CFLAGS=-mmacosx-version-min=10.13`,
// 	// 	`CGO_CXXFLAGS=-I/Users/klyuchnikovv/go/src/github.com/KlyuchnikovV/archive/edicode/build`,
// 	// 	`CGO_ENABLED=1`,
// 	// 	`CGO_LDFLAGS=-framework UniformTypeIdentifiers -mmacosx-version-min=10.13`,
// 	// 	`GOOS=darwin`,
// 	// 	`GOARCH=arm64`,
// 	// }
// 	cmd.Env = append(cmd.Env, os.Environ()...)
// 	cmd.Stdout = &outBuffer
// 	cmd.Stderr = &errBuffer

// 	err := cmd.Run()
// 	if outBuffer.Len() > 0 {
// 		fmt.Printf("build plugin frontend: %s", outBuffer.String())
// 	}

// 	if err == nil {
// 		return nil
// 	}

// 	if outBuffer.Len() > 0 {
// 		fmt.Printf("build plugin frontend: %s", outBuffer.String())
// 	}

// 	return fmt.Errorf(
// 		"can't build plugin frontend: %s: %s", errBuffer.String(), err,
// 	)
// }

func (plug *Plugin) LoadLib(core Context) error {
	plugin, err := plugin.Open(
		path.Join(plug.Directory(), plug.Executable()),
	)
	if err != nil {
		return err
	}

	plug.plugin, err = plugin.Lookup("Plugin")
	if err != nil {
		return err
	}

	backend, ok := plug.plugin.(BackendPlugin)
	if !ok {
		// TODO: errors package
		return fmt.Errorf("plugin has wrong type")
	}

	if err := backend.Init(core); err != nil {
		return err
	}

	plug.Functions = backend.Functions

	return nil
}

func (plug *Plugin) Call(f string, args interface{}) (interface{}, error) {
	return plug.Functions()[f](args)
}

func (plug *Plugin) MakeAction(f string) error {
	if plug.Actions == nil {
		panic("not implemented")
		return nil
	}

	action, ok := plug.Actions()[f]
	if !ok {
		panic("not action")
	}

	var err = action()

	return err
}
