// Package main contains plugins.
package main

import "fmt"

// Command is a command plugin symbol name.
var Command command //nolint:deadcode,gochecknoglobals // ...

type command struct{}

// Execute method executes plugin logic.
func (c command) Execute() error {
	fmt.Println("Command: Hello World!") //nolint:forbidigo // ...

	return nil
}
