package dev

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/common"
	"github.com/thedodd/barge/core"
	"github.com/thedodd/barge/drivers"
	"github.com/thedodd/barge/registry"
	"github.com/thedodd/barge/testutils"
)

func setUp(data []byte) (tmpDir string, config *core.Bargefile, cmd *UpCommand, ui *cli.MockUi, cb func()) {
	// Build a *UpCommand instance.
	ui = &cli.MockUi{}
	cmd = &UpCommand{ui}

	// Create a temporary directory for a test to run.
	tmpDir, _ = ioutil.TempDir("/tmp", "barge")
	originalWD, _ := os.Getwd()
	os.Chdir(tmpDir)

	// Write the given Bargefile data.
	if data != nil {
		ioutil.WriteFile("Bargefile", data, 0777)
		config, _ = common.GetConfig(cmd.UI)
	} else {
		config = &core.Bargefile{Development: &core.DevEnvConfig{}}
	}

	return tmpDir, config, cmd, ui, func() {
		os.Chdir(originalWD)
		os.RemoveAll(tmpDir)
	}
}

func patchRegistry() func() {
	// Patch the registry.
	origRegistry := registry.Registry
	newRegistry := make(map[string]core.Driver)
	for key := range registry.Registry {
		newRegistry[key] = &testutils.MockDriver{}
	}
	registry.Registry = newRegistry

	return func() {
		registry.Registry = origRegistry
	}
}

////////////////////
// Tests for Run. //
////////////////////
func TestRunHandlesErrorWhereBargefileIsInvalid(t *testing.T) {
	_, _, cmd, ui, cleanup := setUp(nil)
	defer cleanup()

	output := cmd.Run([]string{})

	if 1 != output {
		t.Errorf("Expected return code 1, got: %d", output)
	}
	if "Bargefile not found in current directory.\n" != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
}

func TestRunReturns0WithSuccess(t *testing.T) {
	_, config, cmd, ui, cleanup := setUp(testutils.DevelopmentBargefile)
	defer cleanup()
	registryCleanup := patchRegistry()
	defer registryCleanup()
	mockedDriver := registry.Registry["virtualbox"].(*testutils.MockDriver)
	mockedDriver.On("Up", config, ui).Return(0)

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

/////////////////////
// Tests for Help. //
/////////////////////
func TestHelpReturnsExpectedText(t *testing.T) {
	cmd := &UpCommand{&cli.MockUi{}}
	expected := "Spin up a docker machine according to the Bargefile's development section."

	output := cmd.Help()

	if expected != output {
		t.Errorf("Unexpected output: %s", output)
	}
}

/////////////////////////
// Tests for Synopsis. //
/////////////////////////
func TestSynopsisReturnsExpectedText(t *testing.T) {
	cmd := &UpCommand{&cli.MockUi{}}
	expected := "Spin up a docker machine."

	output := cmd.Synopsis()

	if expected != output {
		t.Errorf("Unexpected output: %s", output)
	}
}

/////////////////////////////
// Tests for selectDriver. //
/////////////////////////////
func TestSelectDriverReturnsExpectedDriver(t *testing.T) {
	_, config, _, ui, cleanup := setUp(testutils.DevelopmentBargefile)
	defer cleanup()
	expectedDriver := registry.Registry[config.Development.Driver]

	driver := selectDriver(config, ui)
	_, ok := driver.(*drivers.VirtualBox)

	if driver != expectedDriver || !ok {
		t.Error("Unexpected driver selected.")
	}
}
