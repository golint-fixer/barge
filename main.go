// barge
// Development and deployment for docker based apps made easy.
// Inspiration for this project is based on the ObjectRocket devtools
// ecosystem and the HashiCorp otto project.
package main

import (
	"os"

	"github.com/mitchellh/cli"

	"github.com/thedodd/barge/dev"
)

const (
	bargeVersion = "Barge 0.0.0"
	name         = "barge"
)

func main() {
	// Build the CLI UI.
	baseUI := &cli.ConcurrentUi{Ui: &cli.BasicUi{Writer: os.Stdout}}
	ui := &cli.ColoredUi{
		OutputColor: cli.UiColorNone,
		InfoColor:   cli.UiColorNone,
		ErrorColor:  cli.UiColorRed,
		WarnColor:   cli.UiColorYellow,
		Ui:          baseUI,
	}

	// Build command factory and register it with the top-level CLI.
	commands := map[string]cli.CommandFactory{
		"dev": func() (cli.Command, error) {
			return &dev.UpCommand{UI: ui}, nil
		},
		// "dev destroy": func() (cli.Command, error) {
		// 	return &dev.DestroyCommand{UI: ui}, nil
		// },
	}
	bargeCLI := &cli.CLI{
		Args:     os.Args[1:],
		Name:     name,
		Commands: commands,
		Version:  bargeVersion,
	}

	// Run the top-level CLI.
	exitStatus, err := bargeCLI.Run()
	if err != nil {
		ui.Error(err.Error())
	}

	os.Exit(exitStatus)
}
