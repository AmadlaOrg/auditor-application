package cmd

import (
	"github.com/AmadlaOrg/auditor-application/amadla"
	"github.com/AmadlaOrg/auditor-application/amadla/entity"
	"github.com/spf13/cobra"
)

var AmadlaCmd = &cobra.Command{
	Use:   "amadla",
	Short: "Amadla commands",
}

func init() {
	AmadlaCmd.AddCommand(entity.ListEntitiesCmd)
	AmadlaCmd.AddCommand(amadla.VersionSupportCmd)
}
