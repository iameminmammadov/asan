package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a command by a tag",
	Long: `The 'get' command returns a command that is associated with the provided tag.
Example:
  asan get -t gcloud`,
	Run: func(cmd *cobra.Command, args []string) {
		if tag == "" {
			fmt.Println("Error: Provide a tag")
			cmd.Usage()
			os.Exit(1)
		}

		data := loadYaml()

		command, exists := data.Entries[tag]
		if !exists {
			fmt.Printf("There is no command associated with the tag: %s", tag)
			os.Exit(1)
		}

		fmt.Println(command)

	},
}

func init() {
	getCmd.Flags().StringVarP(&tag, "tag", "t", "", "Tag for the command (required)")
	getCmd.MarkFlagRequired("tag")
	rootCmd.AddCommand(getCmd)
}
