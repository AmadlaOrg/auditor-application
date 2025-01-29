package cmd

import (
	amadlaCmd "github.com/AmadlaOrg/auditor-application/amadla/cmd"
	entityCmd "github.com/AmadlaOrg/auditor-application/amadla/entity/cmd"
	"github.com/spf13/cobra"
)

var AmadlaCmd = &cobra.Command{
	Use:   "amadla",
	Short: "Amadla commands",
}

func init() {
	AmadlaCmd.AddCommand(entityCmd.ListEntitiesCmd)
	AmadlaCmd.AddCommand(amadlaCmd.VersionSupportCmd)
}
