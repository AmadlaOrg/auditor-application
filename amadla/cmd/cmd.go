package cmd

import (
	"github.com/AmadlaOrg/auditor-application/amadla"
	"github.com/spf13/cobra"
)

var VersionSupportCmd = &cobra.Command{
	Use:   "version",
	Short: "Supported version",
	Run:   runVersionSupportCmd,
}

// runVersionSupportCmd
func runVersionSupportCmd(cmd *cobra.Command, args []string) {
	amadlaService := amadla.NewAmadlaService()
	cmd.Println(amadlaService.VersionSupport())
}
