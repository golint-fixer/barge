package drivers

import (
	"fmt"
	"os/exec"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/core"
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

// Up spin up a docker machine according to the Bargefile specs.
func (vb *VirtualBox) Up(bargefile *core.Bargefile, ui cli.Ui) int {
	// Build up the command to execute.
	uiWriter := &cli.UiWriter{Ui: ui}
	cmd := exec.Command(
		"docker-machine",
		"create", "--driver", "virtualbox",
		"--virtualbox-cpu-count", fmt.Sprint(bargefile.Development.CPUS),
		"--virtualbox-disk-size", fmt.Sprint(bargefile.Development.Disk),
		"--virtualbox-memory", fmt.Sprint(bargefile.Development.RAM),
		bargefile.Development.MachineName,
	)
	cmd.Stdout = uiWriter
	cmd.Stderr = uiWriter

	// Wrap command for testability.
	cmdIface := core.CommandWrapper(cmd)

	// Execute command and expose any errors.
	if err := cmdIface.Run(); err != nil {
		ui.Error(err.Error())
		return 1
	}
	return 0
}

// Destroy the docker machine.
func (vb *VirtualBox) Destroy(bargefile *core.Bargefile, ui cli.Ui) int {
	return 0
}

// Rebuild the docker machine.
func (vb *VirtualBox) Rebuild(bargefile *core.Bargefile, ui cli.Ui) int {
	return 0
}
