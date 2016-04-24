package common

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/thedodd/barge/testutils"
)

func setUp(data []byte) (tmpDir string, ui *cli.MockUi, cb func()) {
	// Build a *Command instance.
	ui = &cli.MockUi{}

	// Create a temporary directory for a test to run.
	tmpDir, _ = ioutil.TempDir("/tmp", "barge")
	originalWD, _ := os.Getwd()
	os.Chdir(tmpDir)

	// Write the given Bargefile data.
	if data != nil {
		ioutil.WriteFile("Bargefile", data, 0777)
	}

	return tmpDir, ui, func() {
		os.Chdir(originalWD)
		os.RemoveAll(tmpDir)
	}
}

//////////////////////////
// Tests for GetConfig. //
//////////////////////////
func TestGetConfigErrsWhereNoBargefileExists(t *testing.T) {
	_, ui, cleanup := setUp(nil)
	defer cleanup()

	_, err := GetConfig(ui)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err.Error() != "Bargefile not found in current directory." {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
}

func TestGetConfigErrsWhereBargefileIsMalformedAndFailsUnmarshaling(t *testing.T) {
	_, ui, cleanup := setUp([]byte("malformed"))
	defer cleanup()

	_, err := GetConfig(ui)

	if err.Error() != "Error reading Bargefile." {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
	if "key 'malformed' expected start of object ('{') or assignment ('=')\n" != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
	if "" != ui.OutputWriter.String() {
		t.Errorf("Unexpected UI output: %s", ui.OutputWriter.String())
	}
}

func TestGetConfigErrsWhereBargefileUnmarshalContainsBadType(t *testing.T) {
	_, ui, cleanup := setUp([]byte(`development {ram = "not a valid int"}`))
	defer cleanup()

	_, err := GetConfig(ui)

	if "Error(s) parsing Bargefile." != err.Error() {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
	if "'development.ram' expected type 'int', got unconvertible type 'string'\n" != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
	if "" != ui.OutputWriter.String() {
		t.Errorf("Unexpected UI output: %s", ui.OutputWriter.String())
	}
}

func TestGetConfigErrsWhereBargefileValidationFailsOnProvider(t *testing.T) {
	_, ui, cleanup := setUp(testutils.DevelopmentInvalidDriver)
	defer cleanup()

	_, err := GetConfig(ui)

	if "Error(s) validating Bargefile." != err.Error() {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
	if "Driver must be one of: [virtualbox] Given value: invalidDriver\n" != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
	if "" != ui.OutputWriter.String() {
		t.Errorf("Unexpected UI output: %s", ui.OutputWriter.String())
	}
}

func TestGetConfigErrsWhereBargefileValidationFailsOnStandardField(t *testing.T) {
	_, ui, cleanup := setUp(testutils.DevelopmentInvalidDisk)
	defer cleanup()

	_, err := GetConfig(ui)

	if "Error(s) validating Bargefile." != err.Error() {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
	if "Validation failed for field 'Development.Disk' with validator 'min=5120' and value '5119'.\n" != ui.ErrorWriter.String() {
		t.Errorf("Unexpected UI error: %s", ui.ErrorWriter.String())
	}
	if "" != ui.OutputWriter.String() {
		t.Errorf("Unexpected UI output: %s", ui.OutputWriter.String())
	}
}

func TestGetConfigReturnsExpectedBargefile(t *testing.T) {
	_, ui, cleanup := setUp(testutils.DevelopmentBargefile)
	defer cleanup()

	bargefile, err := GetConfig(ui)
	dev := bargefile.Development

	if err != nil {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
	if dev.Disk != 5120 || dev.MachineName != "devVM" || dev.Network != "bridge" || dev.Driver != "virtualbox" || dev.RAM != 1024 {
		t.Errorf("Bargefile.Development not populated with expected values: %+v", bargefile)
	}
}
