package dev

import (
	"fmt"

	"github.com/thedodd/barge/common"
	"github.com/thedodd/barge/core"
	"github.com/thedodd/barge/registry"

	"github.com/mitchellh/cli"
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

	// Select the driver to use for development.
	driver := selectDriver(config, cmd.UI)
	fmt.Println(fmt.Sprintf("%T: %+v", driver, driver))

	return 0
}

// Synopsis of the `dev` command.
func (cmd *Command) Synopsis() string {
	return "Synopsis of `dev` command."
}

func selectDriver(config *core.Bargefile, ui cli.Ui) core.Driver {
	// Validation of allowed drivers is taken core of by the configuration system.
	// No need to validate here.
	return registry.Registry[config.Development.Driver]
}
