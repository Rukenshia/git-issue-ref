package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"regexp"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var file string

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply the format to a commit message",
	Long: `Applies the ref format to the given commit message. Tries to either get the ref from
the format (if no '{ref}' is specified), the commit message beginning (ABC-123 message)
or last but not least the branch (i.e. feature/ABC-123-something).

By default, this command will fail if no reference could be found unless link was provided with the --non-intrusive flag.`,
	Run: func(cmd *cobra.Command, args []string) {
		str := strings.Join(args, " ")

		if len(file) > 0 {
			f, err := ioutil.ReadFile(file)
			if err != nil {
				color.Red("invalid commit message file")
				os.Exit(1)
			}

			str = string(f)
		}

		ref, err := parseRef(str)

		if checkFormat(format) && err != nil {
			if err.Error() == "no issue ref found" {
				color.Yellow("\ngit-issue-ref WARN: no issue ref found, adding nothing to commit message\n\n")
			} else {
				color.Red("\ngit-issue-ref FAILED: %s\n\n", err.Error())
				os.Exit(1)
			}
		}

		var output = str
		if len(ref) > 0 {
			output = fmt.Sprintf("%s%s",
				strings.Replace(format, "{ref}", ref, -1),
				strings.TrimSpace(strings.Replace(str, ref, "", -1)))
		}

		if len(file) > 0 {
			if err := ioutil.WriteFile(file, []byte(output), 0644); err != nil {
				color.Red("\ngit-issue-ref FAILED: %s\n\n", err.Error())
				os.Exit(1)
			}
		} else {
			color.Green(output)
		}
	},
}

func init() {
	RootCmd.AddCommand(applyCmd)

	applyCmd.Flags().StringVar(&format, "format", "[{ref}] ", "format of the commit message prefix")
	applyCmd.Flags().StringVar(&file, "file", "", "read and write to this file instead of stdin/out")
	applyCmd.Flags().BoolVar(&nonIntrusive, "non-intrusive", false, "if no reference could be found, omit it and dont fail")
}

var reRef = regexp.MustCompile(`^([\w-]*?-?[0-9]{1,})`)
var reBranchRef = regexp.MustCompile(`^(.*?\/)?([\w-]*?-?[0-9]{1,})`)

func parseRef(str string) (string, error) {
	m := reRef.FindStringSubmatch(str)

	if len(m) == 0 || len(m[1]) == 0 {
		return inferFromBranch()
	}

	return m[1], nil
}

func inferFromBranch() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")

	branch, err := cmd.CombinedOutput()
	if err != nil {
		if err.Error() == "exit status 128" {
			msg := string(branch)
			if strings.Contains(msg, "ambiguous argument") {
				return "", errors.New("no issue ref found")
			}
			return "", errors.New(string(branch))
		}
		return "", err
	}

	m := reBranchRef.FindStringSubmatch(string(branch))
	if len(m) == 0 || len(m[1]) == 0 {
		return "", errors.New("no issue ref found")
	}

	return m[2], nil
}
