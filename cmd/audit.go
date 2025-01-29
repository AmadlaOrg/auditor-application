package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"os"
)

var (
	// Cmd flags package global variable
	entityPath string

	// Functions
	osOpen = os.Open

	// AuditCmd sets up the audit command
	AuditCmd = &cobra.Command{
		Use:   "audit",
		Short: "Audit commands",
		Run:   runAuditCmd,
	}
)

// runAuditCmd
func runAuditCmd(cmd *cobra.Command, args []string) {
	// 1. Setup of the `entity` flags
	cmd.Flags().StringVarP(
		&entityPath,
		"entity",
		"e",
		"",
		"Specify the entity file path (optional)",
	)

	// 2. The entity flag is required
	err := cmd.MarkFlagRequired("entity")
	if err != nil {
		cmd.Println(err)
		return
	}

	var entity io.Reader

	// 3. Check if the entity path or stdin is available
	if entityPath == "" {
		// If no `--entity` flag is provided, use stdin
		stat, err := os.Stdin.Stat()
		if err != nil {
			cmd.Printf("Failed to stat stdin: %v\n", err)
			return
		}
		// Check if stdin has data (not a terminal input)
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			entity = os.Stdin
		} else {
			cmd.Println("No entity file path or stdin input provided")
			if err := cmd.Execute(); err != nil {
				cmd.Println(err)
			}
			return
		}
	} else {
		// Open the entity file if the `--entity` flag is provided
		entityFile, err := osOpen(entityPath)
		if err != nil {
			if os.IsNotExist(err) {
				cmd.Println("Entity file does not exist")
			} else {
				cmd.Printf("Failed to open entity file: %v\n", err)
			}
			return
		}
		defer func(file *os.File) {
			if err := file.Close(); err != nil {
				cmd.Printf("Failed to close entity file: %v\n", err)
			}
		}(entityFile)
		entity = entityFile
	}

	// Use the `entity` io.Reader as needed
	cmd.Println("Entity successfully loaded:", entity)
}
