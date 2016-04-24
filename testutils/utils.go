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
	DevelopmentBargefile     = []byte("development {\ndisk = 5120\nmachineName = \"devVM\"\nnetwork = \"bridge\"\ndriver = \"virtualbox\"\nram = 1024}")
	DevelopmentInvalidDriver = []byte("development {\ndisk = 5120\nmachineName = \"devVM\"\nnetwork = \"bridge\"\ndriver = \"invalidDriver\"\nram = 1024}")
	DevelopmentInvalidDisk   = []byte("development {\ndisk = 5119\nmachineName = \"devVM\"\nnetwork = \"bridge\"\ndriver = \"virtualbox\"\nram = 1024}")
)

///////////////////////////////
// Driver registry fixtures. //
///////////////////////////////

// MockDriver a mock implementing the core.Driver interface for testing.
type MockDriver struct {
	mock.Mock
}

// Deps return a slice of *Dep describing all needed command-line tools for this driver.
func (driver *MockDriver) Deps() []*core.Dep {
	driver.Called()
	return []*core.Dep{}
}

// Start a docker machine according to the Bargefile using this driver.
func (driver *MockDriver) Start(bargefile *core.Bargefile, ui cli.Ui) int {
	driver.Called(bargefile, ui)
	return 0
}

// Stop the docker machine running under this driver for the given Bargefile.
func (driver *MockDriver) Stop(bargefile *core.Bargefile, ui cli.Ui) int {
	driver.Called(bargefile, ui)
	return 0
}

// Restart the docker machine running under this driver for the given Bargefile.
func (driver *MockDriver) Restart(bargefile *core.Bargefile, ui cli.Ui) int {
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
