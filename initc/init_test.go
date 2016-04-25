package initc

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/testutils"
)

func setUp(data []byte) (tmpDir string, cmd *InitCommand, ui *cli.MockUi, cb func()) {
	// Build a *UpCommand instance.
	ui = &cli.MockUi{}
	cmd = &InitCommand{ui}

	// Create a temporary directory for a test to run.
	tmpDir, _ = ioutil.TempDir("/tmp", "barge")
	originalWD, _ := os.Getwd()
	os.Chdir(tmpDir)

	return tmpDir, cmd, ui, func() {
		os.Chdir(originalWD)
		os.RemoveAll(tmpDir)
	}
}

////////////////////
// Tests for Run. //
////////////////////
func TestRunReturns0WithSuccess(t *testing.T) {
	_, cmd, ui, cleanup := setUp(testutils.DevelopmentBargefile)
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
	cmd := &InitCommand{&cli.MockUi{}}
	expected := "Initialize a Bargefile."

	output := cmd.Help()

	if expected != output {
		t.Errorf("Unexpected output: %s", output)
	}
}

/////////////////////////
// Tests for Synopsis. //
/////////////////////////
func TestSynopsisReturnsExpectedText(t *testing.T) {
	cmd := &InitCommand{&cli.MockUi{}}
	expected := "Initialize a Bargefile."

	output := cmd.Synopsis()

	if expected != output {
		t.Errorf("Unexpected output: %s", output)
	}
}
