package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

var (
	// Functions
	jsonMarshal = json.Marshal
	yamlMarshal = yaml.Marshal

	// AmadlaCmd
	AmadlaCmd = &cobra.Command{
		Use:   "amadla",
		Short: "Amadla commands",
		Run:   runAmadla,
	}
)

func init() {
	// Add flags to the command (Cobra automatically handles --help)
	AmadlaCmd.Flags().BoolP("json", "j", false, "Display output in JSON format")
	AmadlaCmd.Flags().BoolP("yaml", "y", false, "Display output in YAML format")
}

func runAmadla(cmd *cobra.Command, args []string) {
	// Retrieve flags
	jsonFlag, _ := cmd.Flags().GetBool("json")
	yamlFlag, _ := cmd.Flags().GetBool("yaml")

	// Error if both --json and --yaml are set
	if jsonFlag && yamlFlag {
		cmd.Println("Error: Cannot use both --json and --yaml flags at the same time.")
		os.Exit(1)
	}

	// Define data
	amadlaSupport := map[string]any{
		"applications": map[string]string{
			"hery":  "^0", // TODO:
			"judge": "^0", // TODO:
		},
		"entities": map[string]string{
			"github.com/AmadlaOrg/EntityApplication": "^v1.0.0",
		},
	}

	// Handle JSON output
	if jsonFlag {
		jsonData, err := jsonMarshal(amadlaSupport)
		if err != nil {
			cmd.Println("Error encoding JSON:", err)
			os.Exit(1)
		}
		cmd.Println(string(jsonData))
		return
	}

	// Handle YAML output
	if yamlFlag {
		yamlData, err := yamlMarshal(amadlaSupport)
		if err != nil {
			cmd.Println("Error encoding YAML:", err)
			os.Exit(1)
		}
		cmd.Println(string(yamlData))
		return
	}

	// Default: Display as table
	displayAsTable(amadlaSupport)
}

// Function to display the data in a table format
func displayAsTable(requirements map[string]any) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Category", "Supported", "Version Supported"})

	for category, data := range requirements {
		switch v := data.(type) {
		case map[string]string:
			for key, value := range v {
				table.Append([]string{category, key, value})
			}
		case map[string]any:
			for key, subData := range v {
				if subMap, ok := subData.(map[string]string); ok {
					for subKey, subValue := range subMap {
						table.Append([]string{category, fmt.Sprintf("%s.%s", key, subKey), subValue})
					}
				}
			}
		}
	}

	table.Render()
}
