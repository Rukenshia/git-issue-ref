package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// unlinkCmd represents the unlink command
var unlinkCmd = &cobra.Command{
	Use:   "unlink",
	Short: "Removes the commit-msg hook",
	Long:  `Removes any active commit-msg hook`,
	Run: func(cmd *cobra.Command, args []string) {
		if s, err := os.Stat(".git/hooks/commit-msg"); err != nil || s.IsDir() {
			color.Red("commit-msg hook does not exist")
			os.Exit(1)
		}

		if err := os.Remove(".git/hooks/commit-msg"); err != nil {
			color.Red("could not remove hook: %s\n", err)
			os.Exit(1)
		}

		color.Green("hook removed")
	},
}

func init() {
	RootCmd.AddCommand(unlinkCmd)
}
