package plugins

import (
	"fmt"
	"strings"
)

func (manager *Manager) Actions(filterBy string) map[string]func() error {
	var actions = make(map[string]func() error)

	for _, plug := range manager.plugins {
		if plug.Actions == nil {
			continue
		}

		for actionName, action := range plug.Actions() {
			actions[fmt.Sprintf("%s: %s", plug.Config.Name, actionName)] = action
		}
	}

	return actions
}

func (manager *Manager) RightPanels() []PanelEntry {
	var rightPanels = make([]PanelEntry, 0, len(manager.plugins))

	for _, plug := range manager.plugins {
		if plug.Config.Panels.RightPanel == nil {
			continue
		}

		rightPanels = append(rightPanels, *plug.Config.Panels.RightPanel)
	}

	return rightPanels
}

func (manager *Manager) Plugins() []*Plugin {
	var plugins = make([]*Plugin, 0, len(manager.plugins))

	for _, plugin := range manager.plugins {
		plugins = append(plugins, plugin)
	}

	return plugins
}

func (manager *Manager) Call(f string, args interface{}) (interface{}, error) {
	var slice = strings.Split(f, ".")

	switch len(slice) {
	case 0:
		return nil, fmt.Errorf("namespace not found")
	case 1:
		return nil, fmt.Errorf("function not found")
	case 2:
		return manager.plugins[slice[0]].Call(slice[1], args)
	default:
		return nil, fmt.Errorf("too much namespaces")
	}
}

func (manager *Manager) MakeAction(f string) (interface{}, error) {
	var slice = strings.Split(f, ".")

	switch len(slice) {
	case 0:
		return nil, fmt.Errorf("namespace not found")
	case 1:
		return nil, fmt.Errorf("function not found")
	case 2:
		return nil, manager.plugins[slice[0]].MakeAction(slice[1])
	default:
		return nil, fmt.Errorf("too much namespaces")
	}
}
