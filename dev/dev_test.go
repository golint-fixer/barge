package dev

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/common"
)

var (
	developmentBargefile = []byte("development {\ndisk = 5120\nmachineName = \"devVM\"\nnetwork = \"bridge\"\ndriver = \"virtualbox\"\nram = 1024}")
)

func setUp(data []byte) (tmpDir string, config *common.Bargefile, cmd *Command, ui *cli.MockUi, cb func()) {
	// Build a *Command instance.
	ui = &cli.MockUi{}
	cmd = &Command{ui}

	// Create a temporary directory for a test to run.
	tmpDir, _ = ioutil.TempDir("/tmp", "barge")
	originalWD, _ := os.Getwd()
	os.Chdir(tmpDir)

	// Write the given Bargefile data.
	if data != nil {
		ioutil.WriteFile("Bargefile", data, 0777)
		config, _ = common.GetConfig(cmd.UI)
	} else {
		config = &common.Bargefile{Development: &common.DevEnvConfig{}}
	}

	return tmpDir, config, cmd, ui, func() {
		os.Chdir(originalWD)
		os.RemoveAll(tmpDir)
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
	_, _, cmd, ui, cleanup := setUp(developmentBargefile)
	defer cleanup()

	output := cmd.Run([]string{})

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
	cmd := &Command{&cli.MockUi{}}
	expected := "Help text for `dev` command."

	output := cmd.Help()

	if expected != output {
		t.Errorf("Unexpected output: %s", output)
	}
}

/////////////////////////
// Tests for Synopsis. //
/////////////////////////
func TestSynopsisReturnsExpectedText(t *testing.T) {
	cmd := &Command{&cli.MockUi{}}
	expected := "Synopsis of `dev` command."

	output := cmd.Synopsis()

	if expected != output {
		t.Errorf("Unexpected output: %s", output)
	}
}
