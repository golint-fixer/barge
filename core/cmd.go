package core

import "os/exec"

// CommandWrapperSig the signature of a function which wraps system commands.
// README(TheDodd): WHAT IS THIS? This allows for the interception and wrapping of commands which
// would be executed against the OS. Test implementations of the signature can return a stub which
// records the commands usage.
type CommandWrapperSig func(*exec.Cmd) CmdInterface

// CommandWrapper the CommandWrapperSig function to use.
var CommandWrapper CommandWrapperSig = wrapCommand

// CmdInterface an interface wrapping the os/exec.Cmd struct. Makes command execution unit testable.
type CmdInterface interface {
	CombinedOutput() ([]byte, error)
	Output() ([]byte, error)
	Run() error
	Start() error
	// StderrPipe() (io.ReadCloser, error)
	// StdinPipe() (io.WriteCloser, error)
	// StdoutPipe() (io.ReadCloser, error)
	Wait() error
}

// A default implementation of the CommandWrapperSig. Does nothing other than return the given
// cmd wrapped in a CmdInterface.
func wrapCommand(cmd *exec.Cmd) CmdInterface {
	return cmd
}
