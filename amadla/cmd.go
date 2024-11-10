package amadla

import "github.com/spf13/cobra"

var VersionSupportCmd = &cobra.Command{
	Use:   "version",
	Short: "Supported version",
	Run: func(cmd *cobra.Command, args []string) {
		amadlaService := NewAmadlaService()
		cmd.Println(amadlaService.VersionSupport())
	},
}
