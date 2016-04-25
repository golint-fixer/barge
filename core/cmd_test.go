package core

import (
	"os/exec"
	"testing"
)

////////////////////////////
// Tests for wrapCommand. //
////////////////////////////
func TestWrapCommandReturnsGivenCommandExactly(t *testing.T) {
	cmd := exec.Command("true", "some", "arguments")

	output := wrapCommand(cmd)

	if output != cmd {
		t.Error("Expected output to be identical to input.")
	}
}
