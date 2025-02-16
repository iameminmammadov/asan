package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type Commands struct {
	Entries map[string]string `yaml:"commands"` // Defines struct field which is map with keys (string) and values (string)
}

var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save a shell command with a tag",
	Long: `The 'save' command stores frequently used shell commands.
Example:
  asan save -t gcloud "gcloud get credentials k8s_cluster --region=eu-west4-b"`,
	Run: func(cmd *cobra.Command, args []string) {
		if tag == "" || len(args) == 0 {
			fmt.Println("Error: Provide a tag and command.")
			cmd.Usage()
			os.Exit(1)
		}

		command := args[0]

		data := loadYaml()

		if existingCommand, exists := data.Entries[tag]; exists {
			fmt.Printf("The tag %s already exists with the command: \n%s\n", tag, existingCommand)
			fmt.Print("\nDo you want to overwrite it? (y/n) ")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			userInput := scanner.Text()

			// Don't want to overwrite it
			if userInput != "y" && userInput != "Y" && userInput != "Yes" && userInput != "YES" {
				fmt.Println("Command has not been saved")
				return
			}

		}

		data.Entries[tag] = command

		saveYaml(data)

		fmt.Println("Command saved under the tag:", tag)
	},
}

func doesFileExist() {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println("Error creating file:", err)
			os.Exit(1)
		}
		file.Close()
	}
}

// Move LoadYaml into a separate `util.go` function
func loadYaml() Commands {
	doesFileExist()

	var data Commands
	data.Entries = make(map[string]string)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error loading YAML:", err)
		os.Exit(1)
	}

	yaml.Unmarshal(file, &data)

	return data
}

func saveYaml(data Commands) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating YAML", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	encoder.Encode(data)

}


func init() {
	saveCmd.Flags().StringVarP(&tag, "tag", "t", "", "Tag for the command (required)")
	saveCmd.MarkFlagRequired("tag")
	rootCmd.AddCommand(saveCmd)
}
