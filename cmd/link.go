package cmd

import (
	"fmt"
	"strings"

	"os"

	"io/ioutil"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var hookTemplate = `#!/bin/sh
git issue-ref apply {{args}} --file "$1"`

var format string
var force bool
var nonIntrusive bool

// linkCmd represents the link command
var linkCmd = &cobra.Command{
	Use:   "link",
	Short: "Enable git-issue-ref for the current repository",
	Long: `Use this command to enable git-issue-ref for the current repository.
This command will install a git hook that will, depending on your settings, automatically
prepend an issue reference or ask you to put one at the beginning of your message.

By default, commit messages where no ref could be found will be denied (exit status 1). You can override
this behavior by providing the --non-intrusive flag.`,

	Run: func(cmd *cobra.Command, args []string) {
		// check for existing hook
		if !checkFormat(format) {
			color.Red("warn: given format does not include '{ref}', no substition will be available")
		}

		if _, err := os.Stat(".git/hooks/commit-msg"); !force && err == nil {
			color.Red(".git/hooks/commit-msg exists. Use -f to overwrite file")
			os.Exit(1)
		}

		hookArgs := []string{fmt.Sprintf("--format \"%s\"", format)}

		if nonIntrusive {
			hookArgs = append(hookArgs, "--non-intrusive")
		}

		err := ioutil.WriteFile("./.git/hooks/commit-msg",
			[]byte(strings.Replace(hookTemplate, "{{args}}",
				strings.Join(hookArgs, " "), -1)), 0755)

		if err != nil {
			color.Red("%s\n", err)
		}

		color.Green("commit-msg hook created")
	},
}

func init() {
	RootCmd.AddCommand(linkCmd)

	linkCmd.Flags().StringVar(&format, "format", "[{ref}] ", "format of the commit message prefix")
	linkCmd.Flags().BoolVarP(&force, "force", "f", false, "force overwriting existing hooks")
	linkCmd.Flags().BoolVar(&nonIntrusive, "non-intrusive", false, "if no reference could be found, omit it and dont fail")
}

func checkFormat(format string) bool {
	return strings.Contains(format, "{ref}")
}
