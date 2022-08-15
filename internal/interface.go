// Package internal contains internal service logic.
package internal

// PluginLoader plugin handler interface.
type PluginLoader interface {
	Load(module string) error
}

// PluginRunner plugin runner interface.
type PluginRunner interface {
	Run() error
}

// PluginLoadHandler plugin handler interface.
type PluginLoadHandler interface {
	PluginLoader
	PluginRunner
}
