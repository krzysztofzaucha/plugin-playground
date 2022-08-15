// Package plugin contains plugins.
package main

import "fmt"

// WebServer is a webServer plugin symbol name.
var WebServer webServer //nolint:deadcode,gochecknoglobals // ...

type webServer struct{}

// Execute method executes plugin logic.
func (w webServer) Execute() error {
	fmt.Println("WebServer: Hello World!") //nolint:forbidigo // ...

	return nil
}
