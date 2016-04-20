package core

import "github.com/mitchellh/cli"

// Driver describes the interface which must be implemented by any type which is to provide a
// docker-machine driver interface for Barge.
type Driver interface {
	// Deps
	// Return a slice of *Dep describing all needed command-line tools for this driver.
	Deps() []*Dep

	// Start
	// Start a docker machine according to the Bargefile using this driver.
	Start(*Bargefile, cli.Ui) int

	// Stop
	// Stop the docker machine running under this driver for the given Bargefile.
	Stop(*Bargefile, cli.Ui) int

	// Restart
	// Restart the docker machine running under this driver for the given Bargefile.
	Restart(*Bargefile, cli.Ui) int
}

// Dep a type describing a needed command-line tool.
type Dep struct {
	Name       string
	MinVersion string
	MaxVersion string
}
