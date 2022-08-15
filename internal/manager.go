package internal

import (
	"plugin"
	"strings"

	"github.com/krzysztofzaucha/plugin-playground/integration"
	"github.com/pkg/errors"
)

// Manager yeah yeah.
type Manager struct {
	module plugin.Symbol
}

// NewManager yeah yeah.
func NewManager() *Manager {
	return &Manager{} //nolint:exhaustivestruct,exhaustruct // ...
}

// Load loads plugin by name.
func (p *Manager) Load(module string) error {
	// ensure module is provided, if not print help
	if module == "" {
		return errors.Errorf("module name cannot be empty")
	}

	p.module = loadPlugin(module)

	return nil
}

// Run handles plugin execution.
func (p *Manager) Run() error {
	executor, ok := p.module.(integration.PluginExecutor)
	if !ok {
		return errors.Errorf("module is not a plugin, unable to execute")
	}

	if e := executor.Execute(); e != nil {
		return errors.Errorf("module execution failed with error")
	}

	return nil
}

func generateSymbolName(module string) string {
	return strings.ReplaceAll(
		strings.ToTitle(strings.ReplaceAll(module, "-", " ")),
		" ", "")
}

func loadPlugin(module string) plugin.Symbol {
	// locate and load the plugin
	plug, err := plugin.Open("bin/" + module + ".so")
	if err != nil {
		panic(err)
	}

	symName, err := plug.Lookup(generateSymbolName(module))
	if err != nil {
		panic(err)
	}

	return symName
}
