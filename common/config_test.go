package common

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
)

func mockUI(t *testing.T) *cli.MockUi {
	return &cli.MockUi{}
}

func writeMalformedBargefile() string {
	data := []byte("malformed")
	ioutil.WriteFile("Bargefile", data, 0777)
	return string(data[:])
}

//////////////////////////
// Tests for GetConfig. //
//////////////////////////
func TestGetConfigErrsWhereNoBargefileExists(t *testing.T) {
	ui := mockUI(t)
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	_, err := GetConfig(ui)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err.Error() != "Bargefile not found in current directory." {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
}

func TestGetConfigErrsWhereBargefileIsMalformedAndFailsUnmarshalling(t *testing.T) {
	ui := mockUI(t)
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeMalformedBargefile()

	_, err := GetConfig(ui)

	if err == nil {
		t.Error("Expected error to be returned.")
	}
	if err.Error() != "Could not read Bargefile: key 'malformed' expected start of object ('{') or assignment ('=')" {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
}
