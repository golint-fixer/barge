package core

// Bargefile - the Bargefile config to drive this CLI.
type Bargefile struct {
	Development *DevEnvConfig `mapstructure:"development"`
}

// DevEnvConfig barge configuration for the development environment.
type DevEnvConfig struct {
	CPUS        int    `mapstructure:"cpus" validate:"required=true,min=1"`
	Disk        int    `mapstructure:"disk" validate:"required=true,min=5120"`
	Driver      string `mapstructure:"driver" validate:"required=true,validDriver"`
	MachineName string `mapstructure:"machineName" validate:"required=true"`
	RAM         int    `mapstructure:"ram" validate:"required=true"`
}
