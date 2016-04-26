package testutils

import (
	"os/exec"

	"github.com/mitchellh/cli"
	"github.com/stretchr/testify/mock"
	"github.com/thedodd/barge/core"
)

/////////////////////////
// Bargefile fixtures. //
/////////////////////////

var (
	// DevelopmentBargefile - a []byte representing a valid Bargefile.Development section.
	DevelopmentBargefile = []byte("development {\ncpus = 1\ndisk = 5120\nmachineName = \"devVM\"\ndriver = \"virtualbox\"\nram = 1024}")

	// DevelopmentInvalidDriver - a []byte representing a Bargefile.Development section with an invalid driver field.
	DevelopmentInvalidDriver = []byte("development {\ncpus = 1\ndisk = 5120\nmachineName = \"devVM\"\ndriver = \"invalidDriver\"\nram = 1024}")

	// DevelopmentInvalidDisk - a []byte representing a Bargefile.Development section with an invalid disk field.
	DevelopmentInvalidDisk = []byte("development {\ncpus = 1\ndisk = 5119\nmachineName = \"devVM\"\ndriver = \"virtualbox\"\nram = 1024}")
)

///////////////////////////////
// Driver registry fixtures. //
///////////////////////////////

// MockDriver a mock implementing the core.Driver interface for testing.
type MockDriver struct {
	mock.Mock
}

// Deps -
func (driver *MockDriver) Deps() []*core.Dep {
	driver.Called()
	return []*core.Dep{}
}

// Up -
func (driver *MockDriver) Up(bargefile *core.Bargefile, ui cli.Ui) int {
	driver.Called(bargefile, ui)
	return 0
}

// Destroy -
func (driver *MockDriver) Destroy(bargefile *core.Bargefile, ui cli.Ui) int {
	driver.Called(bargefile, ui)
	return 0
}

// Rebuild -
func (driver *MockDriver) Rebuild(bargefile *core.Bargefile, ui cli.Ui) int {
	driver.Called(bargefile, ui)
	return 0
}

////////////////////////////////
// exec.Cmd related fixtures. //
////////////////////////////////

// CommandWrapper for testing to stand in for the core.cmd.CommandWrapper.
func CommandWrapper(mock *MockCmd) core.CommandWrapperSig {
	return func(cmd *exec.Cmd) core.CmdInterface {
		mock.MockedCmd = cmd
		return mock
	}
}

// MockCmd a mock implementing the core.CmdInterface for testing.
type MockCmd struct {
	mock.Mock
	MockedCmd *exec.Cmd
}

// CombinedOutput -
func (cmd *MockCmd) CombinedOutput() ([]byte, error) {
	cmd.Called()
	return []byte(""), nil
}

// Output -
func (cmd *MockCmd) Output() ([]byte, error) {
	cmd.Called()
	return []byte(""), nil
}

// Run -
func (cmd *MockCmd) Run() error {
	args := cmd.Called()
	return args.Error(0)
}

// Start -
func (cmd *MockCmd) Start() error {
	cmd.Called()
	return nil
}

// Wait -
func (cmd *MockCmd) Wait() error {
	cmd.Called()
	return nil
}
