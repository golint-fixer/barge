// barge
// Development and deployment for docker based apps made easy.
// Inspiration for this project is based on the ObjectRocket devtools
// ecosystem and the HashiCorp otto project.
package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"

	"github.com/thedodd/barge/commands"
)

// Commands is the mapping of all the available barge commands.
var Commands map[string]cli.CommandFactory

func init() {
	ui := &cli.BasicUi{Writer: os.Stdout}

	Commands = map[string]cli.CommandFactory{
		"dev": func() (cli.Command, error) {
			return &commands.DevCommand{UI: ui}, nil
		},
	}
}

func main() {
	bargeCLI := cli.NewCLI("barge", "0.0.0")
	bargeCLI.Args = os.Args[1:]
	bargeCLI.Commands = Commands

	exitStatus, err := bargeCLI.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
