package plugins

import (
	"os"
)

type Manager struct {
	ctx     Context
	config  Config
	plugins map[string]*Plugin
}

func New(ctx Context, configPath string) (*Manager, error) {
	var config = new(Config)
	if err := config.Load(configPath); err != nil {
		return nil, err
	}

	// var (
	// 	err     error
	// 	plugins = make(map[string]*Plugin, len(config.Plugins))
	// )

	// for _, name := range config.Plugins {
	// 	plugins[name], err = NewPlugin(ctx, path.Join(config.PluginsPath, name, name))
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	return &Manager{
		ctx:    ctx,
		config: *config,
		// plugins: plugins,
	}, nil
}

func (manager *Manager) LoadPlugins() error {
	entries, err := os.ReadDir(manager.config.PluginsPath)
	if err != nil {
		return err
	}

	manager.plugins = make(map[string]*Plugin, len(manager.config.Plugins))

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if !manager.config.Plugins.Has(entry.Name()) {
			// TODO: mark as disabled
			continue
		}

		manager.plugins[entry.Name()], err = NewPlugin(
			manager.ctx,
			manager.config.PluginsPath,
			manager.config.BuildPath,
			entry.Name(),
		)
		if err != nil {
			return err
		}
	}

	manager.ctx.Emit("plugins_loaded", "all", nil)
	return nil

}
