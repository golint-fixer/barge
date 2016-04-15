package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/hashicorp/hcl"
	"github.com/mitchellh/mapstructure"
)

// Bargefile - the Bargefile config to drive this CLI.
type Bargefile struct {
	Development *DevEnvConfig `mapstructure:"development"`
}

// DevEnvConfig barge configuration for the development environment.
type DevEnvConfig struct {
	Disk        int    `mapstructure:"disk"`
	MachineName string `mapstructure:"machineName"`
	Network     string `mapstructure:"network"`
	RAM         int    `mapstructure:"ram"`
}

// GetConfig taken from the local Bargefile. Will panic if Bargefile not found or if there was trouble reading the file.
func GetConfig() (*Bargefile, error) {
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
	rawBargefile := &map[string]map[string]interface{}{}
	if err := hcl.Unmarshal(bargeBytes, rawBargefile); err != nil {
		return nil, fmt.Errorf("Could not read Bargefile: %s", err)
	}

	// Map raw Bargefile config onto Bargefile struct.
	fmt.Println(rawBargefile)
	bargefile := &Bargefile{&DevEnvConfig{}}
	if err := mapstructure.Decode(rawBargefile, bargefile); err != nil {
		return nil, fmt.Errorf("Error processing Bargefile: %s", err)
	}
	return bargefile, nil
}
