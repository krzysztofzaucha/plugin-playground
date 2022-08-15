// Package integration contains interfaces for plugins integration.
package integration

// PluginExecutor describes execution interface.
type PluginExecutor interface {
	Execute() error
}
