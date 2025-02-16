package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a command by a tag",
	Long: `The 'run' runs a command that is associated with the provided tag.
Example:
  asan run -t gcloud`,
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

		runShellCommand(command)

	},
}

func runShellCommand(command string) {

	cmd := exec.Command("sh", "-c", command)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = cmd.Wait()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			fmt.Printf("Command failed with exit status %d\n", exitError.ExitCode())
		} else {
			fmt.Printf("Error executing command: %v\n", err)
		}
		os.Exit(1)
	}

}

func init() {
	runCmd.Flags().StringVarP(&tag, "tag", "t", "", "Tag for the command (required)")
	runCmd.MarkFlagRequired("tag")
	rootCmd.AddCommand(runCmd)
}
