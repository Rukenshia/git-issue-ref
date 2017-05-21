package main

import (
	"os"

	"github.com/Rukenshia/git-issue-ref/cmd"

	"github.com/fatih/color"
)

func main() {
	if s, err := os.Stat(".git"); err != nil || !s.IsDir() {
		color.Red("fatal: Not a git repository (or any of the parent directories): .git")
		return
	}

	cmd.Execute()
}
