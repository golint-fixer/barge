package dev

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
)

var (
	developmentBargefile = []byte("development {\ndisk = 5120\nmachineName = \"devVM\"\nnetwork = \"bridge\"\nprovider = \"virtualbox\"\nram = 1024}")
)

// mockUI - a mock cli.Ui for testing.
func mockUI(t *testing.T) *cli.MockUi {
	return &cli.MockUi{}
}

// writeBargefile - write the given []byte to ./Bargefile.
func writeBargefile(data []byte) string {
	ioutil.WriteFile("Bargefile", data, 0777)
	return string(data[:])
}

////////////////////
// Tests for Run. //
////////////////////
func TestRunHandlesErrorWhereBargefileIsInvalid(t *testing.T) {
	ui := mockUI(t)
	cmd := &Command{ui}

	output := cmd.Run([]string{})

	if 1 != output {
		t.Errorf("Expected return code 1, got: %d", output)
	}
	if "Bargefile not found in current directory.\n" != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
}

func TestRunReturns0WithSuccess(t *testing.T) {
	ui := mockUI(t)
	cmd := &Command{ui}
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile(developmentBargefile)

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
