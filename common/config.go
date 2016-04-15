package common

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/hashicorp/hcl"
	"github.com/mitchellh/cli"
	"github.com/mitchellh/mapstructure"
	"gopkg.in/go-playground/validator.v8"
)

var (
	vldtr          = validator.New(&validator.Config{TagName: "validate"})
	validProviders = []string{"virtualbox"}
)

func init() {
	vldtr.RegisterValidation("validProvider", validProvider)
}

// Bargefile - the Bargefile config to drive this CLI.
type Bargefile struct {
	Development *DevEnvConfig `mapstructure:"development"`
}

// DevEnvConfig barge configuration for the development environment.
type DevEnvConfig struct {
	Disk        int    `mapstructure:"disk" validate:"required=true,min=5120"`
	MachineName string `mapstructure:"machineName" validate:"required=true"`
	Network     string `mapstructure:"network" validate:"required=true"`
	Provider    string `mapstructure:"provider" validate:"required=true,validProvider"`
	RAM         int    `mapstructure:"ram" validate:"required=true"`
}

// GetConfig taken from the local Bargefile. Will panic if Bargefile not found or if there was trouble reading the file.
func GetConfig(ui cli.Ui) (*Bargefile, error) {
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
	bargefile := &Bargefile{&DevEnvConfig{}}
	if err := mapstructure.Decode(rawBargefile, bargefile); err != nil {
		return nil, fmt.Errorf("Error processing Bargefile: %s", err)
	}

	// Validate Bargefile.
	if err := vldtr.Struct(bargefile); err != nil {
		errs, _ := err.(validator.ValidationErrors)
		handleValidationErrors(ui, bargefile, errs)
		return nil, errors.New("Errors validating Bargefile.")
	}

	return bargefile, nil
}

func handleValidationErrors(ui cli.Ui, bargefile *Bargefile, errs validator.ValidationErrors) {
	for _, fieldError := range errs {
		switch fieldError.Tag {

		case "validProvider":
			ui.Error(fmt.Sprintf("Provider must be one of: %s", validProviders))

		default:
			ui.Error(
				fmt.Sprintf(
					"Validation failed for field '%s' with validator '%s=%s' and value '%s'.",
					fieldError.NameNamespace,
					fieldError.Tag,
					fieldError.Param,
					fieldError.Value,
				),
			)
		}
	}
}

func validProvider(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {
	valid := false
	for _, val := range validProviders {
		if param == val {
			valid = true
		}
	}
	return valid
}
