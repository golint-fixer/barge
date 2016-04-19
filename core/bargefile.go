package core

// Bargefile - the Bargefile config to drive this CLI.
type Bargefile struct {
	Development *DevEnvConfig `mapstructure:"development"`
}

// DevEnvConfig barge configuration for the development environment.
type DevEnvConfig struct {
	Disk        int    `mapstructure:"disk" validate:"required=true,min=5120"`
	MachineName string `mapstructure:"machineName" validate:"required=true"`
	Network     string `mapstructure:"network" validate:"required=true"`
	Driver      string `mapstructure:"driver" validate:"required=true,validDriver"`
	RAM         int    `mapstructure:"ram" validate:"required=true"`
}
