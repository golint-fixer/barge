package drivers

import (
	"github.com/thedodd/barge/core"

	"github.com/mitchellh/cli"
)

// VirtualBox is the Barge driver implementation for interfacing with VirtualBox via docker-machine.
type VirtualBox struct{}

// Deps returns a slice of *Dep describing all needed command-line tools for this driver.
func (vb *VirtualBox) Deps() []*core.Dep {
	return []*core.Dep{
		&core.Dep{Name: "docker", MinVersion: "1.11.0", MaxVersion: ""},
		&core.Dep{Name: "docker-machine", MinVersion: "0.7.0", MaxVersion: ""},
		&core.Dep{Name: "VBoxManage", MinVersion: "5.0.0", MaxVersion: ""},
	}
}

// Start a docker machine according to the Bargefile specs.
func (vb *VirtualBox) Start(bargefile *core.Bargefile, ui cli.Ui) {
	return
}

// Stop the docker machine.
func (vb *VirtualBox) Stop(bargefile *core.Bargefile, ui cli.Ui) {
	return
}

// Restart the docker machine.
func (vb *VirtualBox) Restart(bargefile *core.Bargefile, ui cli.Ui) {
	return
}
