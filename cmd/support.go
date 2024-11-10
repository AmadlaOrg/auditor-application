package cmd

import (
	"github.com/spf13/cobra"
)

var SupportCmd = &cobra.Command{
	Use:   "support",
	Short: "List all the supported entities",
	Run: func(cmd *cobra.Command, args []string) {
	},
}
