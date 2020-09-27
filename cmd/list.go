package cmd

import (
	"github.com/j1fig/staccato/midi"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists available MIDI devices.",
	Run: func(cmd *cobra.Command, args []string) {
		midi.List()
	},
}
