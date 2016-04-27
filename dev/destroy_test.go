package dev_test

import (
	"testing"

	"github.com/thedodd/barge/dev"
)

///////////////////////////////////
// Tests for DestroyCommand.Run. //
///////////////////////////////////
func TestRunReturns0OnSeccess(t *testing.T) {
	cmd := &dev.DestroyCommand{}

	output := cmd.Run([]string{})

	if 0 != output {
		t.Errorf("Expected a return code of `0`, got `%d`.", output)
	}
}

////////////////////////////////////
// Tests for DestroyCommand.Help. //
////////////////////////////////////
func TestHelpReturnsExpectedString(t *testing.T) {
	cmd := &dev.DestroyCommand{}

	output := cmd.Help()

	if "Destroy the docker machine defined in this project's Bargefile." != output {
		t.Error("Unexpected help message.")
	}
}

////////////////////////////////////////
// Tests for DestroyCommand.Synopsis. //
////////////////////////////////////////
func TestSynopsisReturnsExpectedString(t *testing.T) {
	cmd := &dev.DestroyCommand{}

	output := cmd.Synopsis()

	if "Destroy the docker machine defined in this project's Bargefile." != output {
		t.Error("Unexpected help message.")
	}
}
