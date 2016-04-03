// Package dev
// The idea for this command is that it will provision a docker-machine for
// your project based on the name of the project (I.E., repo name) or the
// project name specified in the config file.
package dev

import "github.com/mitchellh/cli"

// Command interface implementation for the `dev` command.
type Command struct {
	UI cli.Ui
}

// Help text for the `dev` command.
func (cmd *Command) Help() string {
	return "Help text for `dev` command."
}

// Run the `dev` command.
func (cmd *Command) Run(args []string) int {
	return 1
}

// Synopsis of the `dev` command.
func (cmd *Command) Synopsis() string {
	return "Synopsis of `dev` command."
}
