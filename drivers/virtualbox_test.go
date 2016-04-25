package drivers

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/core"
	"github.com/thedodd/barge/testutils"
)

func setUp() (*core.Bargefile, *cli.MockUi, *VirtualBox, *testutils.MockCmd, func()) {
	bargefile := &core.Bargefile{
		Development: &core.DevEnvConfig{
			CPUS:        1,
			Disk:        5120,
			Driver:      "virtualbox",
			MachineName: "test",
			RAM:         1024,
		},
	}
	ui := &cli.MockUi{}
	vb := &VirtualBox{}

	// NOTICE: CommandWrapper is overwritten here for testing purposes.
	origWrapper := core.CommandWrapper
	cmdMock := &testutils.MockCmd{}
	core.CommandWrapper = testutils.CommandWrapper(cmdMock)

	return bargefile, ui, vb, cmdMock, func() {
		core.CommandWrapper = origWrapper
	}
}

////////////////////////////////
// Tests for VirtualBox.Deps. //
////////////////////////////////
func TestDepsReturnsExpectedSlice(t *testing.T) {
	_, _, vb, _, cleanup := setUp()
	defer cleanup()
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

//////////////////////////////
// Tests for VirtualBox.Up. //
//////////////////////////////
func TestUpExecutesExpectedOSCall(t *testing.T) {
	config, ui, vb, cmdMock, cleanup := setUp()
	defer cleanup()
	expectedArgs := []string{
		"docker-machine", "create", "--driver", "virtualbox",
		"--virtualbox-disk-size", fmt.Sprint(config.Development.Disk),
		"--virtualbox-memory", fmt.Sprint(config.Development.RAM),
		config.Development.MachineName,
	}
	cmdMock.On("Run").Return(nil)

	output := vb.Up(config, ui)

	if output != 0 {
		t.Errorf("Expected return code of `%d`, got `%d`.", 0, output)
	}
	if !cmdMock.AssertExpectations(t) {
		t.Error("Unmet expectations.")
	}
	for idx, val := range expectedArgs {
		argVal := cmdMock.MockedCmd.Args[idx]
		if argVal != val {
			t.Errorf("Expected arg `%s` at cmd.Args[%d]; got `%s`.", val, idx, argVal)
		}
	}
}

func TestUpReturns1WithErrDuringOSCall(t *testing.T) {
	config, ui, vb, cmdMock, cleanup := setUp()
	defer cleanup()
	uiError := "Testing error handling."
	cmdMock.On("Run").Return(errors.New(uiError))

	output := vb.Up(config, ui)

	if output != 1 {
		t.Errorf("Expected return code of `%d`, got `%d`.", 1, output)
	}
	if !cmdMock.AssertExpectations(t) {
		t.Error("Unmet expectations.")
	}
	if fmt.Sprintf("%s\n", uiError) != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
}

///////////////////////////////////
// Tests for VirtualBox.Destroy. //
///////////////////////////////////
func TestDestroy(t *testing.T) {
	config, ui, vb, _, cleanup := setUp()
	defer cleanup()

	vb.Destroy(config, ui)
}

///////////////////////////////////
// Tests for VirtualBox.Rebuild. //
///////////////////////////////////
func TestRebuild(t *testing.T) {
	config, ui, vb, _, cleanup := setUp()
	defer cleanup()

	vb.Rebuild(config, ui)
}
