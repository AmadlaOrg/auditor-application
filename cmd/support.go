package cmd

import (
	"github.com/AmadlaOrg/auditor-application/support"
	"github.com/spf13/cobra"
)

var SupportCmd = &cobra.Command{
	Use:   "support",
	Short: "List all the supported entities",
	Run: func(cmd *cobra.Command, args []string) {
		supportService := support.NewSupportService()
		for _, entity := range supportService.List() {
			cmd.Println(entity)
		}
	},
}
