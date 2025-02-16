package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// type Commands struct {
// 	Entries map[string]string `yaml:"commands"` // Defines struct field which is map with keys (string) and values (string)
// }

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all shell commands that have been saved",
	Long: `List all shell commands that have been saved.
Example:
  asan list`,
	Run: func(cmd *cobra.Command, args []string) {
		printAllCommands()
	},
}

func printAllCommands() {
	data := loadYaml()

	if len(data.Entries) == 0 {
		fmt.Println("No commands have been saved. Use `save` to save them.")
		return
	}

	for tag, command := range data.Entries {
		fmt.Printf("%s: %s", tag, command)
	}

}


func init() {
	rootCmd.AddCommand(listCmd)
}
