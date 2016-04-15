// dev - the idea for this command is that it will provision a docker-machine for
// your project based on the name of the project (I.E., repo name) or the
// project name specified in the config file.
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

// Run the `dev` command.
func (cmd *Command) Run(args []string) int {
	// Get runtime config from Bargefile.
	config, err := common.GetConfig()
	if err != nil {
		cmd.UI.Error(err.Error())
		return 1
	}

	// fmt.Println(config)
	fmt.Println(config.Development)
	cmd.UI.Info(fmt.Sprint(config.Development.Disk))
	cmd.UI.Info(config.Development.MachineName)
	cmd.UI.Info(config.Development.Network)
	cmd.UI.Info(fmt.Sprint(config.Development.RAM))
	return 0
}

// Synopsis of the `dev` command.
func (cmd *Command) Synopsis() string {
	return "Synopsis of `dev` command."
}
