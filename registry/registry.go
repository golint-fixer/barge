package registry

import (
	"github.com/thedodd/barge/core"
	"github.com/thedodd/barge/drivers"
)

var (
	// Registry of all available drivers.
	Registry map[string]core.Driver

	// ValidDrivers is a slice of all valid drivers currently registered with Barge.
	ValidDrivers []string
)

func init() {
	// Register all drivers here.
	Registry = map[string]core.Driver{
		"virtualbox": &drivers.VirtualBox{},
	}

	// Populate slice from registered drivers.
	ValidDrivers = make([]string, 0, len(Registry))
	for key := range Registry {
		ValidDrivers = append(ValidDrivers, key)
	}
}
