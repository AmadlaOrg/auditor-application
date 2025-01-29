package cmd

import (
	"github.com/AmadlaOrg/auditor-application/amadla/entity"
	"github.com/spf13/cobra"
)

var ListEntitiesCmd = &cobra.Command{
	Use:   "entities",
	Short: "List all the supported entities",
	Run:   runListEntitiesCmd,
}

// runListEntitiesCmd
func runListEntitiesCmd(cmd *cobra.Command, args []string) {
	entityService := entity.NewEntityService()
	for _, entityList := range entityService.List() {
		cmd.Println(entityList)
	}
}
