// Package cmd contains all commands available for this service
package cmd

import (
	"os"

	"github.com/krzysztofzaucha/plugin-playground/internal/boot"
	"github.com/spf13/cobra"
)

var (
	module    string          //nolint:gochecknoglobals // false positive
	container *boot.Container //nolint:gochecknoglobals // false positive
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{ //nolint:gochecknoglobals,exhaustivestruct,exhaustruct // false positive
	Use:   "service",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

//nolint:gochecknoinits // false positive
func init() {
	// --module=<NAME_OF_A_PLUGIN>
	rootCmd.PersistentFlags().StringVar(
		&module,
		"module",
		"",
		"name of a module to load (empty by default)",
	)

	container = boot.New()
}
