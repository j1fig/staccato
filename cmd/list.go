package cmd

import (
	"github.com/spf13/cobra"
	// "github.com/j1fig/staccato/midi"
	"log"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists available MIDI devices.",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("the list")
	},
}
