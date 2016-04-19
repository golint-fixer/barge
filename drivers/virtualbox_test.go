package drivers

import (
	"testing"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/core"
)

func setUp() (*core.Bargefile, *cli.MockUi, *VirtualBox) {
	bargefile := &core.Bargefile{
		Development: &core.DevEnvConfig{
			Disk:        5120,
			MachineName: "test",
			Network:     "bridge",
			Driver:      "virtualbox",
			RAM:         1024,
		},
	}
	ui := &cli.MockUi{}
	vb := &VirtualBox{}
	return bargefile, ui, vb
}

////////////////////////////////
// Tests for VirtualBox.Deps. //
////////////////////////////////
func TestDepsReturnsExpectedSlice(t *testing.T) {
	_, _, vb := setUp()
	expectedSlice := []*core.Dep{
		&core.Dep{Name: "docker", MinVersion: "1.11.0", MaxVersion: ""},
		&core.Dep{Name: "docker-machine", MinVersion: "0.7.0", MaxVersion: ""},
		&core.Dep{Name: "VBoxManage", MinVersion: "5.0.0", MaxVersion: ""},
	}

	output := vb.Deps()

	for idx, val := range expectedSlice {
		dep := output[idx]
		if dep.Name != val.Name || dep.MinVersion != val.MinVersion || dep.MaxVersion != val.MaxVersion {
			t.Errorf("Unexpected dep: %+v expected: %+v.", dep, val)
		}
	}
}

/////////////////////////////////
// Tests for VirtualBox.Start. //
/////////////////////////////////
func TestStart(t *testing.T) {
	config, ui, vb := setUp()

	vb.Start(config, ui)
}

////////////////////////////////
// Tests for VirtualBox.Stop. //
////////////////////////////////
func TestStop(t *testing.T) {
	config, ui, vb := setUp()

	vb.Stop(config, ui)
}

///////////////////////////////////
// Tests for VirtualBox.Restart. //
///////////////////////////////////
func TestRestart(t *testing.T) {
	config, ui, vb := setUp()

	vb.Restart(config, ui)
}
