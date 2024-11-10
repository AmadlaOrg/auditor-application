package entity

import "github.com/spf13/cobra"

var ListEntitiesCmd = &cobra.Command{
	Use:   "entities",
	Short: "List all the supported entities",
	Run: func(cmd *cobra.Command, args []string) {
		entityService := NewEntityService()
		for _, entity := range entityService.List() {
			cmd.Println(entity)
		}
	},
}
