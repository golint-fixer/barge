package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
)

var config *map[string]interface{}

// Bargefile - the Bargefile config to drive this CLI.
type Bargefile struct {
	Development *DevEnvConfig `mapstructure:"development,squash"`
}

// DevEnvConfig barge configuration for the development environment.
type DevEnvConfig struct {
	Disk        int    `mapstructure:"disk,squash"`
	MachineName string `mapstructure:"machineName,squash"`
	Network     string `mapstructure:"network,squash"`
	RAM         int    `mapstructure:"ram,squash"`
}

// GetConfig taken from the local Bargefile. Will panic if Bargefile not found or if there was trouble reading the file.
func GetConfig() (*map[string]interface{}, error) {
	// Don't parse config multiple times.
	if config != nil {
		return config, nil
	}

	// TODO(TheDodd): make this optional at some point, like the Appfile.
	// If the Bargefile does not exist, then abort.
	if _, err := os.Stat("Bargefile"); err != nil {
		return nil, errors.New("Bargefile not found in current directory.")
	}

	// Read bytes from Bargefile.
	bargeBytes, err := ioutil.ReadFile("Bargefile")
	if err != nil {
		return nil, fmt.Errorf("Error reading Bargefile: %s", err)

	}

	// Unmarshal bytes onto raw map.
	rawBargefile := &map[string]interface{}{}
	if err := hcl.Unmarshal(bargeBytes, rawBargefile); err != nil {
		return nil, fmt.Errorf("Could not read Bargefile: %s", err)
	}
	return rawBargefile, nil
}
