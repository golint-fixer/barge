package dev

import (
	"fmt"

	"github.com/mitchellh/cli"

	"github.com/thedodd/barge/common"
)

// Command interface implementation for the `dev` command.
type Command struct {
	UI cli.Ui
}

// Help text for the `dev` command.
func (cmd *Command) Help() string {
	return "Help text for `dev` command."
}

// Run - the idea for this command is that it will provision a docker-machine for
// your project based on the Bargefile configuration.
func (cmd *Command) Run(args []string) int {
	// Get runtime config from Bargefile.
	config, err := common.GetConfig(cmd.UI)
	if err != nil {
		cmd.UI.Error(err.Error())
		return 1
	}
	fmt.Println(fmt.Sprintf("%+v", config))
	return 0
}

// Synopsis of the `dev` command.
func (cmd *Command) Synopsis() string {
	return "Synopsis of `dev` command."
}
