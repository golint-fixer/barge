package dev

import (
	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/common"
)

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
	// Get runtime config from Bargefile.
	config, err := common.GetConfig(cmd.UI)
	if err != nil {
		cmd.UI.Error(err.Error())
		return 1
	}

	// Select the driver to use for development.
	driver := SelectDriver(config, cmd.UI)

	// TODO(TheDodd): get this logic in line.
	// Ensure driver's dependencies are installed and ready to rock.
	// ensureDeps(driver, config, ui)

	// Execute the drivers `Destroy` method.
	return driver.Destroy(config, cmd.UI)
}

// Synopsis of the `dev destroy` subcommand.
func (cmd *DestroyCommand) Synopsis() string {
	return "Destroy the docker machine defined in this project's Bargefile."
}
