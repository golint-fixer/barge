package dev

import "github.com/mitchellh/cli"

// DestroyCommand interface implementation for the `dev destroy` subcommand.
type DestroyCommand struct {
	UI cli.Ui
}

// Help text for the `dev destroy` subcommand.
func (cmd *DestroyCommand) Help() string {
	return "Destroy the docker machine defined in this project's Bargefile."
}

// Run the stop subcommand to tear down
func (cmd *DestroyCommand) Run(args []string) int {
	return 0
}

// Synopsis of the `dev destroy` subcommand.
func (cmd *DestroyCommand) Synopsis() string {
	return "Destroy the docker machine defined in this project's Bargefile."
}
