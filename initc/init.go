package initc

import "github.com/mitchellh/cli"

// InitCommand command interface implementation for the `init` command.
type InitCommand struct {
	UI cli.Ui
}

// Help text for the `init` command.
func (cmd *InitCommand) Help() string {
	return "Initialize a Bargefile."
}

// Run - initialize a Bargefile for the working directory.
func (cmd *InitCommand) Run(args []string) int {
	return 0
}

// Synopsis of the `init` command.
func (cmd *InitCommand) Synopsis() string {
	return "Initialize a Bargefile."
}
