package cmd

import (
	"github.com/j1fig/staccato/midi"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(streamCmd)
}

var listCmd = &cobra.Command{
	Use:   "stream",
	Short: "Streams MIDI messages from selected port to stdout.",
	Run: func(cmd *cobra.Command, args []string) {
		midi.Stream()
	},
}
