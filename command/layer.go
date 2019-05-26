package command

import (
	"github.com/spf13/cobra"
)

var layers []*cobra.Command

func layerCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "layer",
		Short: "Layer run specify layer plugin with config",
		Long:  `Layer run specify layer plugin with config`,
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	// add every layer command
	for _, c := range layers {
		println("add one layer command")
		cmd.AddCommand(c)
	}
	return cmd
}

// RegisterLayer add one lyaer command to list
func RegisterLayer(c *cobra.Command) {
	layers = append(layers, c)
}
