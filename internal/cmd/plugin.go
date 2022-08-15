package cmd

import (
	"github.com/spf13/cobra"
)

// plugin represents the plugin command.
//nolint:gochecknoglobals,exhaustivestruct,exhaustruct // false positive
var plugin = &cobra.Command{
	Use:   "plugin",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		p := container.Manager()

		err := p.Load(module)
		if err != nil {
			panic(err)
		}

		err = p.Run()
		if err != nil {
			panic(err)
		}
	},
}

//nolint:gochecknoinits // false positive
func init() {
	rootCmd.AddCommand(plugin)
}
