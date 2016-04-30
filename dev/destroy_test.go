package dev_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/config"
	"github.com/thedodd/barge/core"
	"github.com/thedodd/barge/dev"
	"github.com/thedodd/barge/registry"
	"github.com/thedodd/barge/testutils"
)

func setupDestroy(data []byte) (tmpDir string, bargefile *core.Bargefile, cmd *dev.DestroyCommand, ui *cli.MockUi, cb func()) {
	// Build a *dev.UpCommand instance.
	ui = &cli.MockUi{}
	cmd = &dev.DestroyCommand{UI: ui}

	// Create a temporary directory for a test to run.
	tmpDir, _ = ioutil.TempDir("/tmp", "barge")
	originalWD, _ := os.Getwd()
	os.Chdir(tmpDir)

	// Write the given Bargefile data.
	if data != nil {
		ioutil.WriteFile("Bargefile", data, 0777)
		bargefile, _ = config.GetConfig(cmd.UI)
	} else {
		bargefile = &core.Bargefile{Development: &core.DevEnvConfig{}}
	}

	return tmpDir, bargefile, cmd, ui, func() {
		os.Chdir(originalWD)
		os.RemoveAll(tmpDir)
	}
}

///////////////////////////////////
// Tests for DestroyCommand.Run. //
///////////////////////////////////
func TestDestroyCommandRunHandlesErrorWhereBargefileIsInvalid(t *testing.T) {
	_, _, cmd, ui, cleanup := setupDestroy(nil)
	defer cleanup()

	output := cmd.Run([]string{})

	if 1 != output {
		t.Errorf("Expected return code 1, got: %d", output)
	}
	if "Bargefile not found in current directory.\n" != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
}

func TestDestroyCommandRunReturns0WithSuccess(t *testing.T) {
	_, bargefile, cmd, ui, cleanup := setupDestroy(testutils.DevelopmentBargefile)
	defer cleanup()
	registryCleanup := PatchRegistry()
	defer registryCleanup()
	mockedDriver := registry.Registry["virtualbox"].(*testutils.MockDriver)
	mockedDriver.On("Destroy", bargefile, ui).Return(0)

	output := cmd.Run([]string{})

	if !mockedDriver.AssertExpectations(t) {
		t.Errorf("Expected assertions to be correct.")
	}
	if 0 != output {
		t.Errorf("Expected return code 0, got: %d", output)
	}
	if nil != ui.ErrorWriter {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
	if nil != ui.OutputWriter {
		t.Errorf("Unexpected UI output: %s", ui.OutputWriter.String())
	}
}

////////////////////////////////////
// Tests for DestroyCommand.Help. //
////////////////////////////////////
func TestHelpReturnsExpectedString(t *testing.T) {
	cmd := &dev.DestroyCommand{}

	output := cmd.Help()

	if "Destroy the docker machine defined in this project's Bargefile." != output {
		t.Error("Unexpected help message.")
	}
}

////////////////////////////////////////
// Tests for DestroyCommand.Synopsis. //
////////////////////////////////////////
func TestSynopsisReturnsExpectedString(t *testing.T) {
	cmd := &dev.DestroyCommand{}

	output := cmd.Synopsis()

	if "Destroy the docker machine defined in this project's Bargefile." != output {
		t.Error("Unexpected help message.")
	}
}
