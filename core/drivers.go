package core

import "github.com/mitchellh/cli"

// Driver describes the interface which must be implemented by any type which is to provide a
// docker-machine driver interface for Barge.
type Driver interface {
	// Deps
	// Return a slice of *Dep describing all needed command-line tools for this driver.
	Deps() []*Dep

	// Up
	// Up sping up a docker machine according to the Bargefile using this driver. Implementations
	// of this interface MUST ENSURE that the machine is up and in a usable state whenever this
	// command is invoked.
	Up(*Bargefile, cli.Ui) int

	// Destroy
	// Destroy the docker machine running under this driver for the given Bargefile.
	Destroy(*Bargefile, cli.Ui) int

	// Rebuild
	// Rebuild the docker machine running under this driver for the given Bargefile.
	Rebuild(*Bargefile, cli.Ui) int
}

// Dep a type describing a needed command-line tool.
type Dep struct {
	Name       string
	MinVersion string
	MaxVersion string
}
