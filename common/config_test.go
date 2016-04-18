package common

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
)

var (
	developmentBargefile       = []byte("development {\ndisk = 5120\nmachineName = \"devVM\"\nnetwork = \"bridge\"\ndriver = \"virtualbox\"\nram = 1024}")
	developmentInvalidProvider = []byte("development {\ndisk = 5120\nmachineName = \"devVM\"\nnetwork = \"bridge\"\ndriver = \"invalidDriver\"\nram = 1024}")
	developmentInvalidDisk     = []byte("development {\ndisk = 5119\nmachineName = \"devVM\"\nnetwork = \"bridge\"\ndriver = \"virtualbox\"\nram = 1024}")
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

//////////////////////////
// Tests for GetConfig. //
//////////////////////////
func TestGetConfigErrsWhereNoBargefileExists(t *testing.T) {
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	_, err := GetConfig(mockUI(t))

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err.Error() != "Bargefile not found in current directory." {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
}

func TestGetConfigErrsWhereBargefileIsMalformedAndFailsUnmarshaling(t *testing.T) {
	ui := mockUI(t)
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile([]byte("malformed"))

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
	ui := mockUI(t)
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile([]byte(`development {ram = "not a valid int"}`))

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
	ui := mockUI(t)
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile(developmentInvalidProvider)

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
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile(developmentInvalidDisk)

	ui := mockUI(t)
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
	ui := mockUI(t)
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile(developmentBargefile)

	bargefile, err := GetConfig(ui)
	dev := bargefile.Development

	if err != nil {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
	if dev.Disk != 5120 || dev.MachineName != "devVM" || dev.Network != "bridge" || dev.Driver != "virtualbox" || dev.RAM != 1024 {
		t.Errorf("Bargefile.Development not populated with expected values: %+v", bargefile)
	}
}
