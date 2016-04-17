package common

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/mapstructure"
)

func mockUI(t *testing.T) *cli.MockUi {
	return &cli.MockUi{}
}

func writeBargefile(data []byte) string {
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

func TestGetConfigErrsWhereBargefileIsMalformedAndFailsUnmarshaling(t *testing.T) {
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile([]byte("malformed"))

	_, err := GetConfig(mockUI(t))

	if err.Error() != "key 'malformed' expected start of object ('{') or assignment ('=')" {
		t.Errorf("Unexpected error message: %s", err.Error())
	}
}

func TestGetConfigErrsWhereBargefileUnmarshalContainsBadType(t *testing.T) {
	ui := mockUI(t)
	dir, _ := ioutil.TempDir("/tmp", "barge")
	defer os.RemoveAll(dir)
	os.Chdir(dir)
	writeBargefile([]byte(`development {ram = "not a valid int"}`))

	_, err := GetConfig(ui)

	eType, ok := err.(*mapstructure.Error)
	if !ok || len(eType.WrappedErrors()) != 1 {
		t.Error("Unexpected error received.")
	}
	for _, er := range eType.WrappedErrors() {
		if er.Error() != "'development.ram' expected type 'int', got unconvertible type 'string'" {
			t.Errorf("Unexpected mapstructure validation error: %s", er.Error())
		}
	}
}
