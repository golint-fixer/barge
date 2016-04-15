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

const bargeVersion = "Barge 0.0.0"

var (
	// Commands is the mapping of all the available barge commands.
	Commands map[string]cli.CommandFactory

	// UI is the shell frontend used for all output.
	UI cli.Ui
)

func init() {
	baseUI := &cli.BasicUi{Writer: os.Stdout}
	UI = &cli.ColoredUi{
		OutputColor: cli.UiColor{Code: 39, Bold: false}, // Default foreground.
		InfoColor:   cli.UiColor{Code: 39, Bold: false}, // Default foreground.
		ErrorColor:  cli.UiColor{Code: 91, Bold: true},  // Red foreground.
		WarnColor:   cli.UiColor{Code: 93, Bold: true},  // Yellow foreground.
		Ui:          baseUI,
	}

	Commands = map[string]cli.CommandFactory{
		"dev": func() (cli.Command, error) {
			return &dev.Command{UI: UI}, nil
		},
	}
}

func main() {
	bargeCLI := cli.NewCLI("barge", bargeVersion)
	bargeCLI.Args = os.Args[1:]
	bargeCLI.Commands = Commands

	exitStatus, err := bargeCLI.Run()
	if err != nil {
		UI.Error(err.Error())
	}

	os.Exit(exitStatus)
}
