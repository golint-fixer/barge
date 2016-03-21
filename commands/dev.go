package commands

import "github.com/mitchellh/cli"

// DevCommand interface implementation for the `dev` command.
type DevCommand struct {
	UI cli.Ui
}

// Help text for the `dev` command.
func (cmd *DevCommand) Help() string {
	return "Help text for `dev` command."
}

// Run the `dev` command.
func (cmd *DevCommand) Run(args []string) int {
	return 1
}

// Synopsis of the `dev` command.
func (cmd *DevCommand) Synopsis() string {
	return "Synopsis of `dev` command."
}
