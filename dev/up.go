package dev

import (
	"github.com/thedodd/barge/config"
	"github.com/thedodd/barge/core"
	"github.com/thedodd/barge/registry"

	"github.com/mitchellh/cli"
)

// UpCommand interface implementation for the `dev` command.
type UpCommand struct {
	UI cli.Ui
}

// Help text for the `dev` command.
func (cmd *UpCommand) Help() string {
	return "Spin up a docker machine according to this project's Bargefile."
}

// Run - the idea for this command is that it will provision a docker-machine for
// your project based on the Bargefile configuration.
func (cmd *UpCommand) Run(args []string) int {
	// Get runtime config from Bargefile.
	bargefile, err := config.GetConfig(cmd.UI)
	if err != nil {
		cmd.UI.Error(err.Error())
		return 1
	}

	// Select the driver to use for development.
	driver := SelectDriver(bargefile, cmd.UI)

	// TODO(TheDodd): get this logic in line.
	// Ensure driver's dependencies are installed and ready to rock.
	// ensureDeps(driver, bargefile, ui)

	// Execute the drivers `Up` method.
	return driver.Up(bargefile, cmd.UI)
}

// Synopsis of the `dev` command.
func (cmd *UpCommand) Synopsis() string {
	return "Spin up a docker machine according to this project's Bargefile."
}

// SelectDriver will select the docker machine driver according to the given Bargefile config.
func SelectDriver(config *core.Bargefile, ui cli.Ui) core.Driver {
	// Validation of allowed drivers is taken core of by the configuration system.
	// No need to validate here.
	return registry.Registry[config.Development.Driver]
}
